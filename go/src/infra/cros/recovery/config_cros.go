// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package recovery

import (
	"io"
	"strings"
)

// CrosRepairConfig provides config for repair cros setup in the lab task.
func CrosRepairConfig() io.Reader {
	return strings.NewReader(createConfiguration([]configPlan{
		{
			name:      "servo",
			body:      servoRepairPlanBody,
			allowFail: true,
		},
		{
			name:      "cros",
			body:      crosRepairPlanBody,
			allowFail: false,
		},
		{
			name:      "chameleon",
			body:      chameleonPlanBody,
			allowFail: true,
		},
		{
			name:      "bluetooth_peer",
			body:      btpeerRepairPlanBody,
			allowFail: true,
		},
		{
			name:      "close",
			body:      crosClosePlanBody,
			allowFail: true,
		},
	}))
}

// CrosDeployConfig provides config for deploy cros setup in the lab task.
func CrosDeployConfig() io.Reader {
	return strings.NewReader(createConfiguration([]configPlan{
		{
			name:      "servo",
			body:      servoRepairPlanBody,
			allowFail: true,
		},
		{
			name:      "cros",
			body:      crosDeployPlanBody,
			allowFail: false,
		},
		{
			name:      "chameleon",
			body:      chameleonPlanBody,
			allowFail: true,
		},
		{
			name:      "bluetooth_peer",
			body:      btpeerRepairPlanBody,
			allowFail: true,
		},
		{
			name:      "close",
			body:      crosClosePlanBody,
			allowFail: true,
		},
	}))
}

// crosClosePlanBody provides plan to close cros repair/deploy tasks.
const crosClosePlanBody = `
"critical_actions": [
],
"actions": {
}
`