# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from google.appengine.ext import ndb

from gae_libs.pipeline_wrapper import BasePipeline
from libs import analysis_status
from model.flake.flake_try_job_data import FlakeTryJobData
from services.flake_failure import flake_try_job


class ProcessFlakeTryJobResultPipeline(BasePipeline):
  """A pipeline for processing a flake try job result."""

  # Arguments number differs from overridden method - pylint: disable=W0221
  def run(self, revision, commit_position, urlsafe_try_job_key,
          urlsafe_flake_analysis_key):
    """Extracts pass rate information and updates flake analysis.

    Args:
      revision (str): The git hash the try job was run against.
      commit_position (int): The commit position corresponding to |revision|.
      urlsafe_try_job_key (str): The urlsafe key to the corresponding try job
          entity.
      urlsafe_flake_analysis_key (str): The urlsafe key for the master flake
          analysis entity to be updated.
    """
    flake_analysis = ndb.Key(urlsafe=urlsafe_flake_analysis_key).get()
    try_job = ndb.Key(urlsafe=urlsafe_try_job_key).get()
    assert flake_analysis
    assert try_job

    assert len(try_job.try_job_ids) == len(try_job.flake_results)
    try_job_id = try_job.try_job_ids[-1]
    assert try_job_id

    try_job_result = try_job.flake_results[-1]
    try_job_data = FlakeTryJobData.Get(try_job_id)
    assert try_job_data

    if not flake_try_job.IsTryJobResultAtRevisionValid(try_job_result,
                                                       revision):
      try_job.status = analysis_status.ERROR
      try_job_data.error = {
          'message': 'Try job does not contain the necessary information',
          'reason': 'Try job does not contain the necessary information',
      }
      try_job.put()
      try_job_data.put()
      return

    result = try_job_result['report']['result'][revision]

    if isinstance(result, basestring):
      # Result is a string 'infra failed'. Try job failed.
      try_job.status = analysis_status.ERROR
      try_job_data.error = {
          'message': 'Try job failed due to infra error',
          'reason': 'Try job failed due to infra error',
      }
      try_job.put()
      try_job_data.put()
      return

    step_name = flake_analysis.canonical_step_name

    if not flake_try_job.IsTryJobResultAtRevisionValidForStep(
        result, step_name):
      try_job.status = analysis_status.ERROR
      try_job_data.error = {
          'message': 'Try job results are not valid',
          'reason': 'Try job results are not vaild',
      }
      try_job.put()
      try_job_data.put()
      return

    try_job.status = analysis_status.COMPLETED
    try_job.put()

    flake_try_job.UpdateAnalysisDataPointsWithTryJobResult(
        flake_analysis, try_job, commit_position, revision)
