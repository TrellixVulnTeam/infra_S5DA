// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cmd

import (
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"go.chromium.org/chromiumos/infra/proto/go/test_platform/skylab_local_state"
	"infra/libs/skylab/inventory"
)

func TestDutInfoToHostInfoConversion(t *testing.T) {
	Convey("When a DUT info is converted to a host info the result is correct.", t, func() {
		board := "dummy_board"
		sku := "dummy_sku"
		osType := inventory.SchedulableLabels_OS_TYPE_CROS
		input := inventory.DeviceUnderTest{
			Common: &inventory.CommonDeviceSpecs{
				Attributes: []*inventory.KeyValue{
					keyValue("sku", "dummy_sku"),
					keyValue("dummy_key", "dummy_value"),
				},
				Labels: &inventory.SchedulableLabels{
					Board:    &board,
					OsType:   &osType,
					Platform: &board,
					Sku:      &sku,
				},
			},
		}

		got := hostInfoFromDutInfo(&input)

		So(got, ShouldNotBeNil)

		// There's no guarantee on the order.
		sort.Strings(got.Labels)

		want := &skylab_local_state.AutotestHostInfo{
			Attributes: map[string]string{
				"dummy_key": "dummy_value",
				"sku":       "dummy_sku",
			},
			Labels: []string{
				"board:dummy_board",
				"device-sku:dummy_sku",
				"os:cros",
				"platform:dummy_board",
			},
			SerializerVersion: 1,
		}

		So(want, ShouldResemble, got)
	})
}

func keyValue(key string, value string) *inventory.KeyValue {
	return &inventory.KeyValue{
		Key:   &key,
		Value: &value,
	}
}
