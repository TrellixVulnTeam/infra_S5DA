// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package algorithms

import configpb "infra/appengine/weetbix/internal/config/proto"

// TestClusteringConfig returns standard clustering configuration
// that can be used for testing.
func TestClusteringConfig() *configpb.Clustering {
	return &configpb.Clustering{}
}
