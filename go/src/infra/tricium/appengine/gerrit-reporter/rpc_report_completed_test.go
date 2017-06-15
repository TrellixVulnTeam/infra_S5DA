// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gerritreporter

import (
	"testing"

	ds "github.com/luci/gae/service/datastore"
	. "github.com/smartystreets/goconvey/convey"

	"golang.org/x/net/context"

	"infra/tricium/api/admin/v1"
	trit "infra/tricium/appengine/common/testing"
	"infra/tricium/appengine/common/track"
)

type mockGerritAPI struct {
	LastMsg      string
	LastComments []*track.Comment
}

func (m *mockGerritAPI) PostReviewMessage(c context.Context, host, change, revision, msg string) error {
	m.LastMsg = msg
	return nil
}
func (m *mockGerritAPI) PostRobotComments(c context.Context, host, change, revision string, runID int64, comments []*track.Comment) error {
	m.LastComments = comments
	return nil
}

func TestReportCompletedRequest(t *testing.T) {
	Convey("Test Environment", t, func() {
		tt := &trit.Testing{}
		ctx := tt.Context()

		request := &track.AnalyzeRequest{
			GitRepo: "https://chromium-review.googlesource.com",
			GitRef:  "refs/changes/88/508788/1",
		}
		So(ds.Put(ctx, request), ShouldBeNil)
		requestKey := ds.KeyForObj(ctx, request)
		run := &track.WorkflowRun{ID: 1, Parent: requestKey}
		So(ds.Put(ctx, run), ShouldBeNil)
		runKey := ds.KeyForObj(ctx, run)
		analyzerName := "Hello"
		So(ds.Put(ctx, &track.AnalyzerRun{
			ID:     analyzerName,
			Parent: runKey,
		}), ShouldBeNil)
		analyzerKey := ds.NewKey(ctx, "AnalyzerRun", analyzerName, 0, runKey)
		So(ds.Put(ctx, &track.AnalyzerRunResult{
			ID:          1,
			Parent:      analyzerKey,
			Name:        analyzerName,
			NumComments: 1,
		}), ShouldBeNil)
		analyzerName = "Lint"
		So(ds.Put(ctx, &track.AnalyzerRun{
			ID:     analyzerName,
			Parent: runKey,
		}), ShouldBeNil)
		analyzerKey = ds.NewKey(ctx, "AnalyzerRun", analyzerName, 0, runKey)
		So(ds.Put(ctx, &track.AnalyzerRunResult{
			ID:          1,
			Parent:      analyzerKey,
			Name:        analyzerName,
			NumComments: 2,
		}), ShouldBeNil)

		Convey("Report completed request", func() {
			mock := &mockGerritAPI{}
			err := reportCompleted(ctx, &admin.ReportCompletedRequest{
				RunId: run.ID,
			}, mock)
			So(err, ShouldBeNil)
			So(mock.LastMsg, ShouldContainSubstring, "3 results")
		})
	})
}
