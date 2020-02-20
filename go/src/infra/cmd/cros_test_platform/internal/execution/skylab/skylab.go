// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package skylab contains the logic for running individual test tasks.
package skylab

import (
	"bytes"
	"context"
	"fmt"
	"infra/cmd/cros_test_platform/internal/execution/isolate"
	"infra/cmd/cros_test_platform/internal/execution/swarming"
	"infra/libs/skylab/request"

	"github.com/golang/protobuf/jsonpb"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/common"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/skylab_test_runner"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/steps"
	swarming_api "go.chromium.org/luci/common/api/swarming/swarming/v1"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/isolated"
	"go.chromium.org/luci/common/logging"
)

// Task represents an individual test task.
type Task struct {
	args          request.Args
	taskReference *TaskReference
	lifeCycle     test_platform.TaskState_LifeCycle
	// Note: If we ever begin supporting other harnesses's result formats
	// then this field will change to a *skylab_test_runner.Result.
	// For now, the autotest-specific variant is more convenient.
	autotestResult *skylab_test_runner.Result_Autotest
}

// NewTask initializes a Task object.
func NewTask(args request.Args) *Task {
	return &Task{args: args}
}

// Name is the task name as it is displayed in the UI.
func (t *Task) Name() string {
	return t.args.Cmd.TaskName
}

// Launch sends an RPC request to start the task.
func (t *Task) Launch(ctx context.Context, c Client) error {
	ref, err := c.LaunchTask(ctx, &t.args)
	if err != nil {
		return errors.Annotate(err, "launch attempt for %s", t.Name()).Err()
	}
	t.taskReference = ref
	t.lifeCycle = test_platform.TaskState_LIFE_CYCLE_PENDING
	logging.Infof(ctx, "Launched attempt for %s as task %s", t.Name(), t.URL())
	return nil
}

// Completed returns whether the current task is complete.
func (t *Task) Completed() bool {
	return t.autotestResult != nil
}

// Verdict aggregates the information about test cases contained in a task into
// a single verdict.
func (t *Task) Verdict() test_platform.TaskState_Verdict {
	if !t.Completed() {
		return test_platform.TaskState_VERDICT_UNSPECIFIED
	}
	if t.autotestResult == nil {
		return test_platform.TaskState_VERDICT_UNSPECIFIED
	}
	if t.autotestResult.Incomplete {
		return test_platform.TaskState_VERDICT_FAILED
	}

	// By default (if no test cases ran), then there is no verdict.
	verdict := test_platform.TaskState_VERDICT_NO_VERDICT
	for _, c := range t.autotestResult.GetTestCases() {
		switch c.Verdict {
		case skylab_test_runner.Result_Autotest_TestCase_VERDICT_FAIL:
			// Any case failing means the flat verdict is a failure.
			return test_platform.TaskState_VERDICT_FAILED
		case skylab_test_runner.Result_Autotest_TestCase_VERDICT_PASS:
			// Otherwise, at least 1 passing verdict means a pass.
			verdict = test_platform.TaskState_VERDICT_PASSED
		case skylab_test_runner.Result_Autotest_TestCase_VERDICT_UNDEFINED:
			// Undefined verdicts do not affect flat verdict.
		}
	}
	return verdict
}

// Tasks with these life cycles contain test results.
// E.g. this excludes killed tasks as they have no chance to produce results.
var lifeCyclesWithResults = map[test_platform.TaskState_LifeCycle]bool{
	test_platform.TaskState_LIFE_CYCLE_COMPLETED: true,
}

// The life cycles that are not final.
var transientLifeCycles = map[test_platform.TaskState_LifeCycle]bool{
	test_platform.TaskState_LIFE_CYCLE_PENDING: true,
	test_platform.TaskState_LIFE_CYCLE_RUNNING: true,
}

// Refresh fetches the latest swarming and isolate state of the given task,
// and updates the task accordingly.
func (t *Task) Refresh(ctx context.Context, clients Client) error {
	results, err := clients.Swarming.GetResults(ctx, []string{t.SwarmingTaskID()})
	if err != nil {
		return errors.Annotate(err, "fetch results").Err()
	}
	result, err := unpackResult(results, t.SwarmingTaskID())
	if err != nil {
		return errors.Annotate(err, "fetch results").Err()
	}

	lc, err := swarming.AsLifeCycle(result.State)
	if err != nil {
		return errors.Annotate(err, "fetch results").Err()
	}
	t.lifeCycle = lc

	switch {
	// Task ran to completion.
	case lifeCyclesWithResults[lc]:
		r, err := getAutotestResult(ctx, result, clients.IsolateGetter)
		if err != nil {
			logging.Debugf(ctx, "failed to fetch autotest results for task %s due to error '%s', treating its results as incomplete (failure)", t.SwarmingTaskID(), err.Error())
			r = &skylab_test_runner.Result_Autotest{Incomplete: true}
		}
		t.autotestResult = r
	// Task no longer running, but didn't run to completion.
	case !transientLifeCycles[lc]:
		t.autotestResult = &skylab_test_runner.Result_Autotest{Incomplete: true}
	// Task is still running.
	default:
	}
	return nil
}

func unpackResult(results []*swarming_api.SwarmingRpcsTaskResult, taskID string) (*swarming_api.SwarmingRpcsTaskResult, error) {
	if len(results) != 1 {
		return nil, errors.Reason("expected 1 result for task id %s, got %d", taskID, len(results)).Err()
	}

	result := results[0]
	if result.TaskId != taskID {
		return nil, errors.Reason("expected result for task id %s, got %s", taskID, result.TaskId).Err()
	}

	return result, nil
}

var liftTestCaseRunnerVerdict = map[skylab_test_runner.Result_Autotest_TestCase_Verdict]test_platform.TaskState_Verdict{
	skylab_test_runner.Result_Autotest_TestCase_VERDICT_PASS: test_platform.TaskState_VERDICT_PASSED,
	skylab_test_runner.Result_Autotest_TestCase_VERDICT_FAIL: test_platform.TaskState_VERDICT_FAILED,
}

// TestCases unpacks test cases contained in the results of a task.
func (t *Task) TestCases() []*steps.ExecuteResponse_TaskResult_TestCaseResult {
	tcs := t.autotestResult.GetTestCases()
	if len(tcs) == 0 {
		// Prefer a nil over an empty slice since it's the proto default.
		return nil
	}
	ret := make([]*steps.ExecuteResponse_TaskResult_TestCaseResult, len(tcs))
	for i, tc := range tcs {
		ret[i] = &steps.ExecuteResponse_TaskResult_TestCaseResult{
			Name:                 tc.Name,
			Verdict:              liftTestCaseRunnerVerdict[tc.Verdict],
			HumanReadableSummary: tc.HumanReadableSummary,
		}
	}
	return ret
}

// URL return the URL of the task page.
func (t *Task) URL() string {
	return t.taskReference.URL()
}

const resultsFileName = "results.json"

func getAutotestResult(ctx context.Context, sResult *swarming_api.SwarmingRpcsTaskResult, gf isolate.GetterFactory) (*skylab_test_runner.Result_Autotest, error) {
	if sResult == nil {
		return nil, errors.Reason("get result: nil swarming result").Err()
	}

	taskID := sResult.TaskId
	outputRef := sResult.OutputsRef
	if outputRef == nil {
		logging.Debugf(ctx, "task %s has no output ref, considering it failed due to incompleteness", taskID)
		return &skylab_test_runner.Result_Autotest{Incomplete: true}, nil
	}

	getter, err := gf(ctx, outputRef.Isolatedserver)
	if err != nil {
		return nil, errors.Annotate(err, "get result").Err()
	}

	logging.Debugf(ctx, "fetching result for task %s from isolate ref %+v", taskID, outputRef)
	content, err := getter.GetFile(ctx, isolated.HexDigest(outputRef.Isolated), resultsFileName)
	if err != nil {
		return nil, errors.Annotate(err, "get result for task %s", taskID).Err()
	}

	var result skylab_test_runner.Result

	err = jsonpb.Unmarshal(bytes.NewReader(content), &result)
	if err != nil {
		return nil, errors.Annotate(err, "get result for task %s", taskID).Err()
	}

	a := result.GetAutotestResult()
	if a == nil {
		return nil, errors.Reason("get result for task %s: no autotest result; other harnesses not yet supported", taskID).Err()
	}

	return a, nil
}

// Result constructs a TaskResults out of the data already contained in the
// Task object. In order to get the latest result, FetchResult needs to be
// called first.
func (t *Task) Result(attemptNum int) *steps.ExecuteResponse_TaskResult {
	logURL := fmt.Sprintf(
		"https://stainless.corp.google.com/browse/chromeos-autotest-results/swarming-%s/",
		t.SwarmingTaskID(),
	)
	gsURL := fmt.Sprintf(
		"gs://chromeos-autotest-results/swarming-%s/",
		t.SwarmingTaskID(),
	)

	return &steps.ExecuteResponse_TaskResult{
		Name: t.Name(),
		State: &test_platform.TaskState{
			LifeCycle: t.lifeCycle,
			Verdict:   t.Verdict(),
		},
		TaskUrl: t.URL(),
		LogUrl:  logURL,
		LogData: &common.TaskLogData{
			GsUrl: gsURL,
		},
		Attempt:   int32(attemptNum),
		TestCases: t.TestCases(),
	}
}

// SwarmingTaskID returns the Swarming task ID.
func (t *Task) SwarmingTaskID() string {
	return t.taskReference.SwarmingTaskID()
}
