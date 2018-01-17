// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gerrit

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	ds "go.chromium.org/gae/service/datastore"

	"infra/tricium/api/admin/v1"
	trit "infra/tricium/appengine/common/testing"
	"infra/tricium/appengine/common/track"
)

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
		functionName := "Hello"
		So(ds.Put(ctx, &track.FunctionRun{
			ID:     functionName,
			Parent: runKey,
		}), ShouldBeNil)
		functionRunKey := ds.NewKey(ctx, "FunctionRun", functionName, 0, runKey)
		So(ds.Put(ctx, &track.FunctionRunResult{
			ID:          1,
			Parent:      functionRunKey,
			Name:        functionName,
			NumComments: 1,
		}), ShouldBeNil)
		functionName = "Lint"
		So(ds.Put(ctx, &track.FunctionRun{
			ID:     functionName,
			Parent: runKey,
		}), ShouldBeNil)
		functionRunKey = ds.NewKey(ctx, "FunctionRun", functionName, 0, runKey)
		So(ds.Put(ctx, &track.FunctionRunResult{
			ID:          1,
			Parent:      functionRunKey,
			Name:        functionName,
			NumComments: 2,
		}), ShouldBeNil)

		Convey("Report completed request", func() {
			mock := &mockRestAPI{}
			err := reportCompleted(ctx, &admin.ReportCompletedRequest{
				RunId: run.ID,
			}, mock)
			So(err, ShouldBeNil)
			So(mock.LastMsg, ShouldContainSubstring, "3 results")
		})

		Convey("Does not report completion when reporting is disabled", func() {
			request.GerritReportingDisabled = true
			So(ds.Put(ctx, request), ShouldBeNil)
			mock := &mockRestAPI{}
			err := reportLaunched(ctx, &admin.ReportLaunchedRequest{
				RunId: run.ID,
			}, mock)
			So(err, ShouldBeNil)
			So(mock.LastMsg, ShouldEqual, "")
		})

	})
}
