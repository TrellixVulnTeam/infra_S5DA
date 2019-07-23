// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package manifest defines structure of YAML files with target definitions.
package manifest

import (
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"go.chromium.org/luci/common/errors"
)

// Manifest is a definition of what to build, how and where.
type Manifest struct {
	// ContextDir is a unix-style path to the docker context directory to ingest
	// (usually a directory with Dockerfile), relative to this YAML file.
	//
	// All symlinks there are resolved to their targets. Only +w and +x file mode
	// bits are preserved. All other file metadata (owners, setuid bits,
	// modification times) are ignored.
	//
	// If not set, the context directory is assumed empty.
	ContextDir string `yaml:"contextdir,omitempty"`

	// Build defines a series of local build steps.
	//
	// Each step may add more files to the context directory. The actual
	// `contextdir` directory on disk won't be modified. Files produced here are
	// stored in a temp directory and the final context directory is constructed
	// from the full recursive copy of `contextdir` and files emitted here.
	Build []*BuildStep `yaml:"build,omitempty"`
}

// BuildStep is one local build operation.
//
// It takes a local checkout and produces one or more output files put into
// the context directory.
//
// This struct is a "case class" with union of all supported build step kinds.
// The chosen "case" is returned by Concrete() method.
type BuildStep struct {
	// Fields common to all build kinds.

	// Dest specifies a location within the context dir to put the result into.
	//
	// Optional in the original YAML, always populated after Manifest is parsed.
	// See individual *BuildStep structs for defaults.
	Dest string `yaml:"dest,omitempty"`

	// Disjoint set of possible build kinds.
	//
	// To add a new step kind:
	//   1. Add a new embedded struct here with definition of the step.
	//   2. Add String() and initStep(...) methods to implement ConcreteBuildStep.
	//   3. Add one more 'if' to initAndSetDefaults(...) below.
	//   4. Add the actual step implementation to builder/step*.go.
	//   5. Add one more type switch to Builder.Build() in builder/builder.go.

	CopyBuildStep `yaml:",inline"` // copy a file or directory into the output
	GoBuildStep   `yaml:",inline"` // build go binary using "go build"

	concrete ConcreteBuildStep // pointer to one of *BuildStep above
}

// ConcreteBuildStep is implemented by various *BuildStep structs.
type ConcreteBuildStep interface {
	String() string // used for human logs only, doesn't have to encode all details

	initStep(bs *BuildStep, cwd string) // populates 'bs' and self
}

// Concrete returns a pointer to some concrete populated *BuildStep.
func (bs *BuildStep) Concrete() ConcreteBuildStep { return bs.concrete }

// CopyBuildStep indicates we want to copy a file or directory.
type CopyBuildStep struct {
	// Copy is a path (relative to the manifest file) to copy files from.
	//
	// Can either be a directory or a file. Whatever it is, it will be put into
	// the output as Dest. By default Dest is a basename of Copy (i.e. we copy
	// Copy into the root of the context dir).
	Copy string `yaml:"copy,omitempty"`
}

func (s *CopyBuildStep) String() string { return fmt.Sprintf("copy %q", s.Copy) }

func (s *CopyBuildStep) initStep(bs *BuildStep, cwd string) {
	normPath(&s.Copy, cwd)
	if bs.Dest == "" {
		bs.Dest = filepath.Base(s.Copy)
	}
}

// GoBuildStep indicates we want to build a go command binary.
type GoBuildStep struct {
	// GoBinary specifies a go command binary to build.
	//
	// This is a path (relative to GOPATH) to some 'main' package. It will be
	// built roughly as:
	//
	//  $ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build <go_binary> -o <dest>
	//
	// Where <dest> (taken from Dest) is relative to the context directory and set
	// to the go package name by default.
	GoBinary string `yaml:"go_binary,omitempty"`
}

func (s *GoBuildStep) String() string { return fmt.Sprintf("go build %q", s.GoBinary) }

func (s *GoBuildStep) initStep(bs *BuildStep, cwd string) {
	if bs.Dest == "" {
		bs.Dest = path.Base(s.GoBinary)
	}
}

// Read reads and initializes the manifest by filling in all defaults.
//
// If cwd is not empty, rebases all relative paths in it on top of it.
func Read(r io.Reader, cwd string) (*Manifest, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Annotate(err, "failed to read the manifest body").Err()
	}
	out := Manifest{}
	if err = yaml.Unmarshal(body, &out); err != nil {
		return nil, errors.Annotate(err, "failed to parse the manifest").Err()
	}
	if err := out.Initialize(cwd); err != nil {
		return nil, err
	}
	return &out, nil
}

// Initialize fills in the defaults.
//
// If cwd is not empty, rebases all relative paths in it on top of it.
//
// Must be called if Manifest{} was allocated in the code (e.g. in unit tests)
// rather than was read via Read(...).
func (m *Manifest) Initialize(cwd string) error {
	normPath(&m.ContextDir, cwd)
	for i := range m.Build {
		if err := initAndSetDefaults(m.Build[i], cwd); err != nil {
			return errors.Annotate(err, "bad build step #%d", i+1).Err()
		}
	}
	return nil
}

func initAndSetDefaults(bs *BuildStep, cwd string) error {
	set := make([]ConcreteBuildStep, 0, 1)
	if bs.CopyBuildStep != (CopyBuildStep{}) {
		set = append(set, &bs.CopyBuildStep)
	}
	if bs.GoBuildStep != (GoBuildStep{}) {
		set = append(set, &bs.GoBuildStep)
	}

	// One and only one substruct should be populated.
	switch {
	case len(set) == 0:
		return errors.Reason("unrecognized or empty").Err()
	case len(set) > 1:
		return errors.Reason("ambiguous").Err()
	default:
		bs.concrete = set[0]
		bs.concrete.initStep(bs, cwd)
		return nil
	}
}

func normPath(p *string, cwd string) {
	if *p != "" {
		*p = filepath.FromSlash(*p)
		if cwd != "" {
			*p = filepath.Join(cwd, *p)
		}
	}
}
