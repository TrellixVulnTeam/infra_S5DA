# Copyright 2015 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import datetime

from google.appengine.ext import ndb
import mock

from components import metrics as metrics_component
from components import utils
from testing_utils import testing
from proto import project_config_pb2
from test import future
import config
import metrics
import model


class MerticsTest(testing.AppengineTestCase):
  def test_send_build_status_metric(self):
    buf = mock.Mock()

    ndb.put_multi([
      model.Build(bucket='chromium', status=model.BuildStatus.SCHEDULED),
      model.Build(bucket='chromium', status=model.BuildStatus.SCHEDULED),
      model.Build(bucket='v8', status=model.BuildStatus.SCHEDULED),
      model.Build(bucket='chromium', status=model.BuildStatus.STARTED),
    ])
    send_future = metrics.send_build_status_metric(
      buf, 'chromium', metrics.METRIC_PENDING_BUILDS,
      model.BuildStatus.SCHEDULED)
    send_future.get_result()
    buf.set_gauge.assert_called_once_with(
      metrics.METRIC_PENDING_BUILDS, 2,
      {metrics.LABEL_BUCKET: 'chromium'})

  def test_send_build_lease_latency(self):
    buf = mock.Mock()
    now = datetime.datetime(2015, 1, 4)
    self.mock(utils, 'utcnow', lambda: now)

    ndb.put_multi([
      model.Build(
        bucket='chromium',
        status=model.BuildStatus.SCHEDULED,
        never_leased=True,
        create_time=datetime.datetime(2015, 1, 1)
      ),
      model.Build(
        bucket='chromium',
        status=model.BuildStatus.SCHEDULED,
        never_leased=True,
        create_time=datetime.datetime(2015, 1, 3)
      ),
      model.Build(
        bucket='chromium',
        status=model.BuildStatus.COMPLETED,
        result=model.BuildResult.CANCELED,
        cancelation_reason=model.CancelationReason.TIMEOUT,
        never_leased=True,
        create_time=datetime.datetime(2015, 1, 3)
      ),
      model.Build(bucket='chromium', status=model.BuildStatus.SCHEDULED),
      model.Build(
        bucket='v8',
        status=model.BuildStatus.SCHEDULED,
        never_leased=True,
        create_time=datetime.datetime(2015, 1, 3)
      ),
    ])
    metrics.send_build_latency(
      buf, metrics.METRIC_LEASE_BUILD_LATENCY, 'chromium', True).get_result()
    buf.set_gauge.assert_called_once_with(
      metrics.METRIC_LEASE_BUILD_LATENCY,
      2.0 * 24 * 3600,  # 2 days,
      {metrics.LABEL_BUCKET: 'chromium'})

  def test_send_all_metrics(self):
    buf = mock.Mock()
    self.mock(metrics_component, 'Buffer', lambda: buf)

    self.mock(config, 'get_buckets_async', mock.Mock())
    config.get_buckets_async.return_value = future([
      project_config_pb2.Bucket(name='x')
    ])
    self.mock(metrics, 'send_build_status_metric', mock.Mock())

    metrics.send_all_metrics()

    metrics.send_build_status_metric.assert_any_call(
      buf, 'x', metrics.METRIC_PENDING_BUILDS, model.BuildStatus.SCHEDULED)
    metrics.send_build_status_metric.assert_any_call(
      buf, 'x', metrics.METRIC_RUNNING_BUILDS, model.BuildStatus.STARTED)
