# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import mock

from model.flake.master_flake_analysis import DataPoint
from model.flake.master_flake_analysis import MasterFlakeAnalysis
from services.flake_failure import confidence_score_util
from waterfall.flake import confidence
from waterfall.flake import flake_constants
from waterfall.test.wf_testcase import WaterfallTestCase


class ConfidenceScoreUtilTest(WaterfallTestCase):

  @mock.patch.object(
      confidence, 'SteppinessForCommitPosition', return_value=0.6)
  def testCalculateCulpritConfidenceScore(self, _):
    analysis = MasterFlakeAnalysis.Create('m', 'b', 124, 's', 't')
    analysis.data_points = [
        DataPoint.Create(pass_rate=0.7, commit_position=1000),
        DataPoint.Create(pass_rate=1.0, commit_position=999)
    ]
    self.assertIsNone(
        confidence_score_util.CalculateCulpritConfidenceScore(analysis, None))
    self.assertEqual(0.6,
                     confidence_score_util.CalculateCulpritConfidenceScore(
                         analysis, 999))

  def testCalculateCulpritConfidenceScoreIntroducedNewFlakyTest(self):
    analysis = MasterFlakeAnalysis.Create('m', 'b', 124, 's', 't')
    analysis.data_points = [
        DataPoint.Create(pass_rate=0.7, commit_position=1000),
        DataPoint.Create(
            pass_rate=flake_constants.PASS_RATE_TEST_NOT_FOUND,
            commit_position=999)
    ]
    self.assertEqual(1.0,
                     confidence_score_util.CalculateCulpritConfidenceScore(
                         analysis, 1000))

  def testCalculateCulpritConfidenceScoreIntroducedStableToFlaky(self):
    analysis = MasterFlakeAnalysis.Create('m', 'b', 124, 's', 't')
    analysis.data_points = [
        DataPoint.Create(pass_rate=0.7, commit_position=1000),
        DataPoint.Create(
            pass_rate=flake_constants.DEFAULT_UPPER_FLAKE_THRESHOLD,
            commit_position=999)
    ]
    self.assertEqual(.7,
                     confidence_score_util.CalculateCulpritConfidenceScore(
                         analysis, 1000))
