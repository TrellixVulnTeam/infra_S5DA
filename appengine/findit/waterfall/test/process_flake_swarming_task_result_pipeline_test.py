# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import datetime
import mock

from infra_api_clients.swarming import swarming_util
from libs import analysis_status
from model.flake.flake_swarming_task import FlakeSwarmingTask
from model.flake.master_flake_analysis import MasterFlakeAnalysis
from services import swarmed_test_util
from waterfall import build_util
from waterfall.build_info import BuildInfo
from waterfall.process_flake_swarming_task_result_pipeline import (
    ProcessFlakeSwarmingTaskResultPipeline)
from waterfall.test import (process_base_swarming_task_result_pipeline_test as
                            base_test)
from waterfall.test import wf_testcase


class ProcessFlakeSwarmingTaskResultPipelineTest(wf_testcase.WaterfallTestCase):

  def _MockedGetSwarmingTaskResultById(self, _host, task_id, _):
    return base_test._SWARMING_TASK_RESULTS[task_id], None

  def setUp(self):
    super(ProcessFlakeSwarmingTaskResultPipelineTest, self).setUp()
    self.pipeline = ProcessFlakeSwarmingTaskResultPipeline()
    self.master_name = 'm'
    self.builder_name = 'b'
    self.build_number = 121
    self.step_name = 'abc_tests on platform'
    self.test_name = 'TestSuite1.test1'
    self.version_number = 1
    self.mock(swarming_util, 'GetSwarmingTaskResultById',
              self._MockedGetSwarmingTaskResultById)

  def testCheckTestsRunStatuses(self):
    analysis = MasterFlakeAnalysis.Create(self.master_name, self.builder_name,
                                          self.build_number, self.step_name,
                                          self.test_name)
    analysis.Save()

    task = FlakeSwarmingTask.Create(self.master_name, self.builder_name,
                                    self.build_number, self.step_name,
                                    self.test_name)
    task.put()

    call_params = ProcessFlakeSwarmingTaskResultPipeline._GetArgs(
        self.pipeline, self.master_name, self.builder_name, self.build_number,
        self.step_name, self.build_number, self.test_name, self.version_number)

    tests_statuses = (
        ProcessFlakeSwarmingTaskResultPipeline._CheckTestsRunStatuses(
            self.pipeline, base_test._SAMPLE_FAILURE_LOG, *call_params))
    self.assertEqual(base_test._EXPECTED_TESTS_STATUS, tests_statuses)

  def testCheckTestsRunStatusesTestExistsNoRunData(self):
    failure_log = {'per_iteration_data': [{}], 'all_tests': [self.test_name]}
    analysis = MasterFlakeAnalysis.Create(self.master_name, self.builder_name,
                                          self.build_number, self.step_name,
                                          self.test_name)
    analysis.Save()

    task = FlakeSwarmingTask.Create(self.master_name, self.builder_name,
                                    self.build_number, self.step_name,
                                    self.test_name)
    task.put()

    call_params = ProcessFlakeSwarmingTaskResultPipeline._GetArgs(
        self.pipeline, self.master_name, self.builder_name, self.build_number,
        self.step_name, self.build_number, self.test_name, self.version_number)

    ProcessFlakeSwarmingTaskResultPipeline._CheckTestsRunStatuses(
        self.pipeline, failure_log, *call_params)

    self.assertEqual(analysis_status.ERROR, task.status)

  @mock.patch.object(
      swarmed_test_util,
      'GetSwarmingTaskFailureLog',
      return_value=(base_test._SAMPLE_FAILURE_LOG, None))
  @mock.patch.object(
      build_util, 'GetBuildInfo', return_value=BuildInfo('m', 'b', 123))
  def testProcessFlakeSwarmingTaskResultPipeline(self, *_):
    # End to end test.
    task = FlakeSwarmingTask.Create(self.master_name, self.builder_name,
                                    self.build_number, self.step_name,
                                    self.test_name)
    task.task_id = 'task_id1'
    task.put()

    analysis = MasterFlakeAnalysis.Create(self.master_name, self.builder_name,
                                          self.build_number, self.step_name,
                                          self.test_name)
    analysis.Save()

    pipeline = ProcessFlakeSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name, 'task_id1', self.build_number, self.test_name,
                 analysis.version_number)
    pipeline.callback(callback_params=pipeline.last_params)
    # Reload from ID to get all internal properties in sync.
    pipeline = ProcessFlakeSwarmingTaskResultPipeline.from_id(
        pipeline.pipeline_id)
    step_name, task_info = pipeline.outputs.default.value
    self.assertEqual('abc_tests', task_info)
    self.assertEqual(self.step_name, step_name)

    task = FlakeSwarmingTask.Get(self.master_name, self.builder_name,
                                 self.build_number, self.step_name,
                                 self.test_name)

    self.assertEqual(analysis_status.COMPLETED, task.status)
    self.assertEqual(base_test._EXPECTED_TESTS_STATUS, task.tests_statuses)

    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 32, 6, 538220), task.created_time)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 32, 9, 90550), task.started_time)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 33, 9), task.completed_time)
