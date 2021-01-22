// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package tasks

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/common/proto"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/server/tq/tqtesting"

	"infra/appengine/rubber-stamper/config"
	"infra/appengine/rubber-stamper/internal/util"
	"infra/appengine/rubber-stamper/tasks/taskspb"
)

func TestQueue(t *testing.T) {
	Convey("Chain works", t, func() {
		cfg := &config.Config{
			HostConfigs: map[string]*config.HostConfig{
				"host": {
					RepoConfigs: map[string]*config.RepoConfig{
						"dummy": {
							BenignFilePattern: &config.BenignFilePattern{
								Paths: []string{"a/b.txt"},
							},
						},
					},
				},
			},
		}

		ctx := memory.Use(context.Background())
		ctx, gerritMock, sched := util.SetupTestingContext(ctx, cfg, "srv-account@example.com", "host", t)

		var succeeded tqtesting.TaskList

		sched.TaskSucceeded = tqtesting.TasksCollector(&succeeded)
		sched.TaskFailed = func(ctx context.Context, task *tqtesting.Task) { panic("should not fail") }

		Convey("Test deduplication", func() {
			host := "host"
			cls := []*gerritpb.ChangeInfo{
				{
					Number:          12345,
					CurrentRevision: "123abc",
					Project:         "dummy",
					Labels: map[string]*gerritpb.LabelInfo{
						"Auto-Submit": {Approved: &gerritpb.AccountInfo{}},
					},
					Revisions: map[string]*gerritpb.RevisionInfo{
						"123abc": {},
					},
				},
				{
					Number:          12345,
					CurrentRevision: "456789",
					Project:         "dummy",
					Revisions: map[string]*gerritpb.RevisionInfo{
						"123abc": {},
					},
				},
			}

			gerritMock.EXPECT().ListFiles(gomock.Any(), proto.MatcherEqual(&gerritpb.ListFilesRequest{
				Number:     cls[0].Number,
				RevisionId: cls[0].CurrentRevision,
			})).Return(&gerritpb.ListFilesResponse{
				Files: map[string]*gerritpb.FileInfo{
					"a/b.txt": nil,
				},
			}, nil)
			gerritMock.EXPECT().ListFiles(gomock.Any(), proto.MatcherEqual(&gerritpb.ListFilesRequest{
				Number:     cls[1].Number,
				RevisionId: cls[1].CurrentRevision,
			})).Return(&gerritpb.ListFilesResponse{
				Files: map[string]*gerritpb.FileInfo{
					"a/b.txt": nil,
				},
			}, nil)

			So(EnqueueChangeReviewTask(ctx, host, cls[0]), ShouldBeNil)
			So(EnqueueChangeReviewTask(ctx, host, cls[0]), ShouldBeNil)
			So(EnqueueChangeReviewTask(ctx, host, cls[1]), ShouldBeNil)
			So(EnqueueChangeReviewTask(ctx, host, cls[1]), ShouldBeNil)
			sched.Run(ctx, tqtesting.StopWhenDrained())
			So(len(succeeded.Payloads()), ShouldEqual, 2)
			So(succeeded.Payloads(), ShouldContain, &taskspb.ChangeReviewTask{
				Host:           "host",
				Number:         12345,
				Revision:       "123abc",
				Repo:           "dummy",
				AutoSubmit:     true,
				RevisionsCount: 1,
			})
			So(succeeded.Payloads(), ShouldContain, &taskspb.ChangeReviewTask{
				Host:           "host",
				Number:         12345,
				Revision:       "456789",
				Repo:           "dummy",
				AutoSubmit:     false,
				RevisionsCount: 1,
			})
		})
	})
}
