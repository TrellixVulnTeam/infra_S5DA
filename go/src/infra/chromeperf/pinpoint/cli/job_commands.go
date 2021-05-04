// Copyright 2021 The Chromium Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"bufio"
	"context"
	"fmt"
	"infra/chromeperf/pinpoint"
	"infra/chromeperf/pinpoint/cli/render"
	"os"
	"strings"

	"github.com/maruel/subcommands"
	"go.chromium.org/luci/common/data/text"
	"go.chromium.org/luci/common/errors"
	"google.golang.org/protobuf/encoding/prototext"
)

type listJobs struct {
	baseCommandRun
	filter string
}

func cmdListJobs(p Param) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: "list-jobs [-filter='']",
		ShortDesc: "Lists jobs tracked by Pinpoint",
		LongDesc: text.Doc(`
			Prints out a list of jobs tracked by Pinpoint to stdout, possibly
			constrained by the filter.
		`),
		CommandRun: wrapCommand(p, func() pinpointCommand {
			return &listJobs{}
		}),
	}
}

func (lj *listJobs) RegisterFlags(p Param) {
	lj.baseCommandRun.RegisterFlags(p)
	// TODO(dberris): Link to documentation about supported fields in the filter.
	lj.Flags.StringVar(&lj.filter, "filter", "", text.Doc(`
		Optional filter to apply to restrict the set of jobs listed. See
		https://aip.dev/160 for details on the filter syntax.
	`))
}

func (lj *listJobs) Run(ctx context.Context, a subcommands.Application, args []string) error {
	c, err := lj.pinpointClient(ctx)
	if err != nil {
		return err
	}

	req := &pinpoint.ListJobsRequest{Filter: lj.filter}
	resp, err := c.ListJobs(ctx, req)
	if err != nil {
		return errors.Annotate(err, "failed during ListJobs").Err()
	}
	if err = render.JobListText(a.GetOut(), resp.Jobs); err != nil {
		return errors.Annotate(err, "failed rendering jobs").Err()
	}
	return nil
}

type getJob struct {
	baseCommandRun
	waitForJobMixin
	downloadResultsMixin
	name string
}

func cmdGetJob(p Param) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: "get-job -name='...'",
		ShortDesc: "Prints information about a Pinpoint job",
		LongDesc: text.Doc(`
			Prints out information about a Job.
			Results can also optionally be downloaded.
		`),
		CommandRun: wrapCommand(p, func() pinpointCommand {
			return &getJob{}
		}),
	}
}

func (gj *getJob) RegisterFlags(p Param) {
	uc := gj.baseCommandRun.RegisterFlags(p)
	gj.downloadResultsMixin.RegisterFlags(&gj.Flags, uc)

	gj.Flags.StringVar(&gj.name, "name", "", text.Doc(`
		Required; the name of the job to get information about.
		Example: "-name=XXXXXXXXXXXXXX"
	`))
}

func (gj *getJob) Run(ctx context.Context, a subcommands.Application, args []string) error {
	if gj.name == "" {
		return errors.Reason("must set -name").Err()
	}

	c, err := gj.pinpointClient(ctx)
	if err != nil {
		return err
	}

	req := &pinpoint.GetJobRequest{Name: pinpoint.LegacyJobName(gj.name)}
	resp, err := c.GetJob(ctx, req)
	if err != nil {
		return errors.Annotate(err, "failed during GetJob").Err()
	}
	out := prototext.MarshalOptions{Multiline: true}.Format(resp)
	fmt.Println(out)

	resp, err = gj.waitForJobMixin.waitForJob(ctx, c, resp, a.GetOut())
	if err != nil {
		return err
	}

	return gj.downloadResultsMixin.doDownloadResults(ctx, resp)
}

type waitJob struct {
	baseCommandRun
	downloadResultsMixin
	quiet bool
	name  string
}

func cmdWaitJob(p Param) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: "wait-job -name='...' [-quiet]",
		ShortDesc: "Polls the Pinpoint backend until the job stops running",
		LongDesc: text.Doc(`
			wait-job blocks until the provided Pinpoint job is finished, i.e.
			until the job is no longer RUNNING nor PENDING. -quiet disables
			informational text output.

			Results can also optionally be downloaded on completion.
		`),
		CommandRun: wrapCommand(p, func() pinpointCommand {
			return &waitJob{}
		}),
	}
}

func (wj *waitJob) RegisterFlags(p Param) {
	uc := wj.baseCommandRun.RegisterFlags(p)
	wj.downloadResultsMixin.RegisterFlags(&wj.Flags, uc)
	wj.Flags.StringVar(&wj.name, "name", "", text.Doc(`
		Required; the name of the job to poll.
		Example: "-name=XXXXXXXXXXXXXX"
	`))
	wj.Flags.BoolVar(&wj.quiet, "quiet", false, text.Doc(`
		Disable informational text output; errors are still printed.
	`))
}

func (wj *waitJob) Run(ctx context.Context, a subcommands.Application, args []string) error {
	if wj.name == "" {
		return errors.Reason("must set -name").Err()
	}

	c, err := wj.pinpointClient(ctx)
	if err != nil {
		return err
	}

	req := &pinpoint.GetJobRequest{Name: pinpoint.LegacyJobName(wj.name)}
	resp, err := c.GetJob(ctx, req)
	if err != nil {
		return errors.Annotate(err, "failed during GetJob").Err()
	}

	// Force `wait` to true because we're always meant to wait with wait-job.
	w := &waitForJobMixin{
		wait:  true,
		quiet: wj.quiet,
	}
	job, err := w.waitForJob(ctx, c, resp, a.GetOut())
	if err != nil {
		return err
	}

	return wj.downloadResultsMixin.doDownloadResults(ctx, job)
}

type cancelJob struct {
	baseCommandRun
	name  string
	force bool
}

func cmdCancelJob(p Param) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: "cancel-job -name='...' 'required reason'",
		ShortDesc: "Cancels an ongoing job",
		LongDesc: text.Doc(`
			Cancels an ongoing job. You must specify a reason as a positional
			argument. Note that you can only cancel jobs you started unless you
			have administrator rights.
		`),
		CommandRun: wrapCommand(p, func() pinpointCommand {
			return &cancelJob{}
		}),
	}
}

func (cj *cancelJob) RegisterFlags(p Param) {
	cj.baseCommandRun.RegisterFlags(p)
	cj.Flags.StringVar(&cj.name, "name", "", text.Doc(`
		Required; the name of the job to Cancel.
		Example: "-name=XXXXXXXXXXXXXX"
	`))
	cj.Flags.BoolVar(&cj.force, "force", false, text.Doc(`
		If unset, the CLI will ask for verification before cancelling the job.
		If set, jobs will be cancelled without additional prompting.
	`))
}

func (cj *cancelJob) Run(ctx context.Context, a subcommands.Application, args []string) error {
	if len(args) != 1 {
		return errors.Reason("Must specify reason as the only positional argument").Err()
	}
	reason := args[0]
	if cj.name == "" {
		return errors.Reason("must set -name").Err()
	}

	c, err := cj.pinpointClient(ctx)
	if err != nil {
		return err
	}

	legacyName := pinpoint.LegacyJobName(cj.name)

	job, err := c.GetJob(ctx, &pinpoint.GetJobRequest{Name: legacyName})
	if err != nil {
		return errors.Annotate(err, "failed during GetJob").Err()
	}
	out := prototext.MarshalOptions{Multiline: true}.Format(job)
	fmt.Println(out)
	fmt.Println("-----------------------------------------")

	if !cj.force {
		fmt.Print("Are you sure you want to cancel the above job? (y/N) ")

		sc := bufio.NewScanner(os.Stdin)
		if !sc.Scan() || !strings.EqualFold(sc.Text(), "y") {
			return errors.Reason("cancelled").Err()
		}
	}

	req := &pinpoint.CancelJobRequest{
		Name:   pinpoint.LegacyJobName(cj.name),
		Reason: reason,
	}
	resp, err := c.CancelJob(ctx, req)
	if err != nil {
		return errors.Annotate(err, "failed during CancelJob").Err()
	}
	out = prototext.MarshalOptions{Multiline: true}.Format(resp)
	fmt.Println(out)
	if ju, err := render.JobURL(job); err != nil {
		fmt.Printf("\nJob cancelled, see %s\n", ju)
	} else {
		fmt.Printf("\nJob cancelled.\n")
	}
	return nil
}
