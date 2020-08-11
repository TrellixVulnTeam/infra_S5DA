# Copyright 2020 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

"""Unit tests for the monitoring module."""

import unittest
import logging
from framework import monitoring

class MonitoringTest(unittest.TestCase):

  def testIncrementAPIRequestsCount(self):
    # Non-service account email gets hidden.
    monitoring.IncrementAPIRequestsCount(
        'v3', 'monorail-prod', client_email='client-email@chicken.com')
    self.assertEqual(
        1,
        monitoring.API_REQUESTS_COUNT.get(
            fields={'client_id': 'monorail-prod',
                    'client_email': 'user@email.com',
                    'version': 'v3'}))

    # None email address gets replaced by 'anonymous'.
    monitoring.IncrementAPIRequestsCount('v3', 'monorail-prod')
    self.assertEqual(
        1,
        monitoring.API_REQUESTS_COUNT.get(
            fields={'client_id': 'monorail-prod', 'client_email': 'anonymous',
                    'version': 'v3'}))

    # Service account email is not hidden
    monitoring.IncrementAPIRequestsCount(
        'endpoints', 'monorail-prod',
        client_email='123456789@developer.gserviceaccount.com')
    self.assertEqual(
        1,
        monitoring.API_REQUESTS_COUNT.get(
            fields={'client_id': 'monorail-prod',
                    'client_email': '123456789@developer.gserviceaccount.com',
                    'version': 'endpoints'}))
