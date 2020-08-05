// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package execution

import (
	"context"
	"fmt"
	"infra/cmd/cros_test_platform/internal/execution/skylab"
	"time"

	"go.chromium.org/chromiumos/infra/proto/go/test_platform"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/config"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/steps"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
)

// RequestTaskSet encapsulates the running state of the set of tasks for one
// cros_test_platform request.
type RequestTaskSet struct {
	// Unique names for tests preserving the order of incoming arguments.
	// This is used to preserve the order in the response.
	testKeys []string

	testTaskSets map[string]*testTaskSet
	retryCounter retryCounter
}

// TaskSetConfig is a wrapper for the parameters common to the testTaskSets.
type TaskSetConfig struct {
	ParentTaskID string
	RequestUID   string
	Deadline     time.Time
}

// NewRequestTaskSet creates a new RequestTaskSet.
func NewRequestTaskSet(tests []*steps.EnumerationResponse_AutotestInvocation, params *test_platform.Request_Params, workerConfig *config.Config_SkylabWorker, tc *TaskSetConfig) (*RequestTaskSet, error) {
	testKeys := make([]string, len(tests))
	testTaskSets := make(map[string]*testTaskSet)
	tm := make(map[string]*steps.EnumerationResponse_AutotestInvocation)
	for i, test := range tests {
		k := testKey(i, test)
		testKeys[i] = k
		t, err := newTestTaskSet(test, params, workerConfig, tc)
		if err != nil {
			return nil, errors.Annotate(err, "new task set").Err()
		}
		testTaskSets[k] = t
		tm[k] = test
	}
	return &RequestTaskSet{
		testKeys:     testKeys,
		testTaskSets: testTaskSets,
		retryCounter: newRetryCounter(params, tm),
	}, nil
}

// testName returns a unique string used as a key for a test.
func testKey(i int, test *steps.EnumerationResponse_AutotestInvocation) string {
	return fmt.Sprintf("%d_%s", i, test.GetTest().GetName())
}

// Completed returns true if all tasks for this request have completed.
func (r *RequestTaskSet) Completed() bool {
	for _, t := range r.testTaskSets {
		if !t.Completed() {
			return false
		}
	}
	return true
}

// LaunchTasks launches initial tasks for all the tests in this request.
func (r *RequestTaskSet) LaunchTasks(ctx context.Context, c skylab.Client) error {
	for _, key := range r.testKeys {
		if err := r.getTestTaskSet(key).LaunchTask(ctx, c); err != nil {
			return err
		}
	}
	return nil
}

func (r *RequestTaskSet) getTestTaskSet(key string) *testTaskSet {
	ts, ok := r.testTaskSets[key]
	if !ok {
		panic(fmt.Sprintf("No test task set for key %s", key))
	}
	return ts
}

// CheckTasksAndRetry checks the status of currently running tasks for this
// request and retries failed tasks when allowed.
func (r *RequestTaskSet) CheckTasksAndRetry(ctx context.Context, c skylab.Client) error {
	for _, key := range r.testKeys {
		ts := r.getTestTaskSet(key)
		if ts.Completed() {
			continue
		}

		latestTask := ts.GetLatestTask()
		rerr := latestTask.Refresh(ctx, c)
		tr := latestTask.Result()
		if rerr != nil {
			return errors.Annotate(rerr, "tick for task %s", tr.LogUrl).Err()
		}

		if !ts.Completed() {
			continue
		}

		logging.Infof(ctx, "Task %s (%s) completed with verdict %s", tr.LogUrl, ts.Name, tr.GetState().GetVerdict())

		shouldRetry, err := r.shouldRetry(ctx, key, latestTask)
		if err != nil {
			return errors.Annotate(err, "tick for task %s", tr.LogUrl).Err()
		}
		if shouldRetry {
			logging.Infof(ctx, "Retrying %s", ts.Name)
			if err := ts.LaunchTask(ctx, c); err != nil {
				return errors.Annotate(err, "tick for task %s: retry test", tr.LogUrl).Err()
			}
			r.retryCounter.NotifyRetry(key)
		}
	}
	return nil
}

// shouldRetry computes if the given test should be retried.
func (r *RequestTaskSet) shouldRetry(ctx context.Context, key string, task *skylab.Task) (bool, error) {
	if !r.retryCounter.CanRetry(ctx, key) {
		return false, nil
	}

	switch v := task.Result().GetState().GetVerdict(); v {
	case test_platform.TaskState_VERDICT_UNSPECIFIED:
		fallthrough
	case test_platform.TaskState_VERDICT_FAILED:
		return true, nil
	case test_platform.TaskState_VERDICT_NO_VERDICT:
		fallthrough
	case test_platform.TaskState_VERDICT_PASSED:
		fallthrough
	case test_platform.TaskState_VERDICT_PASSED_ON_RETRY:
		return false, nil
	default:
		return false, errors.Reason("shouldRetry: unknown verdict %s", v.String()).Err()
	}
}

func (r *RequestTaskSet) response(running bool) *steps.ExecuteResponse {
	resp := &steps.ExecuteResponse{
		TaskResults:         r.taskResults(),
		ConsolidatedResults: r.results(),
		State: &test_platform.TaskState{
			Verdict:   r.verdict(),
			LifeCycle: r.lifecycle(running),
		},
	}
	return resp
}

func (r *RequestTaskSet) lifecycle(running bool) test_platform.TaskState_LifeCycle {
	switch {
	case r.Completed():
		return test_platform.TaskState_LIFE_CYCLE_COMPLETED
	case running:
		return test_platform.TaskState_LIFE_CYCLE_RUNNING
	default:
		// TODO(akeshet): The task set is neither running nor complete, so it
		// was cancelled due to an error while in flight. It's not clear yet
		// if this is the right lifecycle mapping for this state.
		return test_platform.TaskState_LIFE_CYCLE_ABORTED
	}
}

func (r *RequestTaskSet) verdict() test_platform.TaskState_Verdict {
	v := test_platform.TaskState_VERDICT_PASSED
	if !r.Completed() {
		v = test_platform.TaskState_VERDICT_UNSPECIFIED
	}
	for _, t := range r.testTaskSets {
		if !successfulVerdict(t.Verdict()) {
			v = test_platform.TaskState_VERDICT_FAILED
			break
		}
	}
	return v
}

func successfulVerdict(v test_platform.TaskState_Verdict) bool {
	switch v {
	case test_platform.TaskState_VERDICT_PASSED,
		test_platform.TaskState_VERDICT_PASSED_ON_RETRY,
		test_platform.TaskState_VERDICT_NO_VERDICT:
		return true
	default:
		return false
	}
}

func (r *RequestTaskSet) results() []*steps.ExecuteResponse_ConsolidatedResult {
	rs := make([]*steps.ExecuteResponse_ConsolidatedResult, len(r.testKeys))
	for i, key := range r.testKeys {
		rs[i] = &steps.ExecuteResponse_ConsolidatedResult{
			Attempts: r.getTestTaskSet(key).TaskResult(),
		}
	}
	return rs
}

func (r *RequestTaskSet) taskResults() []*steps.ExecuteResponse_TaskResult {
	results := r.results()
	var trs []*steps.ExecuteResponse_TaskResult
	for _, result := range results {
		trs = append(trs, result.Attempts...)
	}
	return trs
}
