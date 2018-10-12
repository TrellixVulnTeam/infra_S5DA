# Copyright 2015 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import json
import logging

from common import constants
from common.findit_http_client import FinditHttpClient
from common.waterfall import buildbucket_client
from common.waterfall import failure_type
from infra_api_clients import logdog_util
from libs import time_util
from model.isolated_target import IsolatedTarget
from model.wf_build import WfBuild
from services import constants as services_constants
from services import swarming
from waterfall import buildbot

HTTP_CLIENT_NO_404_ERROR = FinditHttpClient(no_error_logging_statuses=[404])


def _BuildDataNeedUpdating(build):
  return (not build.data or (
      not build.completed and
      (time_util.GetUTCNow() - build.last_crawled_time).total_seconds() >= 300))


def _GetLogLocationForBuild(build_data):
  """Gets log location for a build.

  For a LUCI build, we can get buildbucket_id of the build from buildbucket
  property and then get log location from buildbucket API.

  For a buildbot build, we can get log location directly from log_location
  property.

  Each property would look like a list like:
    [property_name, property_content, annotation]
  """
  if not build_data:
    return None

  data_json = json.loads(build_data)
  properties = data_json.get('properties') or []
  buildbucket = None
  for property_item in properties:
    if property_item[0] == 'log_location':
      return property_item[1]

    if property_item[0] == 'buildbucket':
      buildbucket = property_item[1]

  if not isinstance(buildbucket, dict):
    return None

  buildbucket_id = buildbucket.get('build', {}).get('id')
  error, build = buildbucket_client.GetTryJobs([buildbucket_id])[0]
  if error:
    logging.exception(
        'Error retrieving buildbucket build id: %s' % buildbucket_id)
    return None

  return logdog_util.GetLogLocationFromBuildbucketBuild(build.response)


def DownloadBuildData(master_name, builder_name, build_number):
  """Downloads build data and returns a WfBuild instance."""
  build = WfBuild.Get(master_name, builder_name, build_number)
  if not build:
    build = WfBuild.Create(master_name, builder_name, build_number)

  status_code = None
  build_updated = False
  # Cache the data to avoid pulling from master again.
  if _BuildDataNeedUpdating(build):
    # Retrieve build data from milo.
    status_code, build.data = buildbot.GetBuildDataFromMilo(
        master_name, builder_name, build_number, HTTP_CLIENT_NO_404_ERROR)
    build.last_crawled_time = time_util.GetUTCNow()
    build_updated = True

  if build.data and not build.log_location:
    build.log_location = _GetLogLocationForBuild(build.data)
    build_updated = True

  if build_updated:
    build.put()

  return status_code, build


def GetBuildInfo(master_name, builder_name, build_number):
  """Gets build info given a master, builder, and build number.

  Args:
    master_name (str): The name of the master.
    builder_name (str): The name of the builder.
    build_number (int): The build number.

  Returns:
    Build information as an instance of BuildInfo.
  """
  status_code, build = DownloadBuildData(master_name, builder_name,
                                         build_number)
  if not build.data:
    return status_code, None
  build_info = buildbot.ExtractBuildInfo(master_name, builder_name,
                                         build_number, build.data)

  if not build.completed:
    build.start_time = build_info.build_start_time
    build.completed = build_info.completed
    build.result = build_info.result
    build.put()

  return status_code, build_info


def GetBuildEndTime(master_name, builder_name, build_number):
  _, build = DownloadBuildData(master_name, builder_name, build_number)
  build_info = buildbot.ExtractBuildInfo(master_name, builder_name,
                                         build_number, build.data)
  return build_info.build_end_time


def GetFailureType(build_info):
  if not build_info.failed_steps:
    return failure_type.UNKNOWN
  # TODO(robertocn): Consider also bailing out of tests with infra failures.
  if constants.COMPILE_STEP_NAME in build_info.failed_steps:
    if build_info.result == buildbot.EXCEPTION:
      return failure_type.INFRA
    return failure_type.COMPILE
  # TODO(http://crbug.com/602733): differentiate test steps from infra ones.
  return failure_type.TEST


def GetLatestBuildNumber(master_name, builder_name):
  """Attempts to get the latest build number on master_name/builder_name."""
  recent_builds = buildbot.GetRecentCompletedBuilds(master_name, builder_name,
                                                    FinditHttpClient())

  if recent_builds is None:
    # Likely a network error.
    logging.error('Failed to detect latest build number on %s, %s', master_name,
                  builder_name)
    return None

  if not recent_builds:
    # In case the builder is new or was recently reset.
    logging.warning('No recent builds found on %s %s', master_name,
                    builder_name)
    return None

  return recent_builds[0]


def GetLatestCommitPosition(master_name, builder_name, target_name):
  """Gets the latest commit position given a master/builder/target."""
  latest_targets = (
      IsolatedTarget.FindLatestIsolateByMaster(
          master_name, builder_name, services_constants.GITILES_HOST,
          services_constants.GITILES_PROJECT, services_constants.GITILES_REF,
          target_name))

  if latest_targets:
    return latest_targets[0].commit_position

  # Fallback to buildbot for builds not yet migrated to LUCI.
  # TODO (crbug.com/804617): Remove fallback logic after migration is complete.
  latest_build_number = GetLatestBuildNumber(master_name, builder_name)

  if not latest_build_number:
    # Something is wrong. Calling code should be responsible for checking for
    # the return value.
    return None

  _, latest_build = GetBuildInfo(master_name, builder_name, latest_build_number)
  return latest_build.commit_position


# TODO(crbug/821865): Remove this after new flake pipelines are stable.
def FindValidBuildNumberForStepNearby(master_name,
                                      builder_name,
                                      step_name,
                                      build_number,
                                      exclude_list=None,
                                      search_distance=3):
  """Finds a valid nearby build number for a step.

  Looks around the given build number for builds that have a reference task
  on swarming. We use this reference swarming task to create a task request,
  and it's required to run the test. If no reference swarming task can be
  found, it's likely that the build failed and the artifact doesn't exist.

  Args:
    master_name (str): Name of the master for this test.
    builder_name (str): Name of the builder for this test.
    step_name (str): Name of the builder for this test.
    build_number (int): Build number to look around.
    exclude_list (lst): Build numbers to exclude from the search.
    search_distance (int): Distance to search on either side of the build.

  Returns:
    (int) Valid nearby build if any, else None.
  """
  builds_to_look_at = [build_number]
  for x in range(1, search_distance + 1):
    builds_to_look_at.append(build_number + x)
    builds_to_look_at.append(build_number - x)

  logging.info('Examining build numbers %r for a valid build',
               builds_to_look_at)

  http_client = FinditHttpClient()
  for build in builds_to_look_at:
    if exclude_list and build in exclude_list:
      continue
    if swarming.CanFindSwarmingTaskFromBuildForAStep(
        http_client, master_name, builder_name, build, step_name):
      return build

  return None


# TODO(crbug/804617): Modify this function to use new LUCI API when it's ready.
def IteratePreviousBuildsFrom(master_name, builder_name, build_number,
                              entry_limit):

  n = build_number - 1
  entry_number = 0
  while n >= 0 and entry_number <= entry_limit:  # pragma: no branch.
    status_code, build_info = GetBuildInfo(master_name, builder_name, n)
    n -= 1
    if build_info:
      entry_number += 1
      yield build_info
    elif status_code == 404:
      continue
    else:
      # 404 means we hit a gap. Otherwise there is something wrong.
      raise Exception('Failed to download build data for build %s/%s/%d' %
                      (master_name, builder_name, n))
