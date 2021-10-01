// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package controller

import (
	"context"
	"time"

	"go.chromium.org/luci/common/logging"

	"infra/cros/hwid"
	ufspb "infra/unifiedfleet/api/v1/models"
	"infra/unifiedfleet/app/model/configuration"
	"infra/unifiedfleet/app/util"
)

const cacheAge = time.Hour

// GetHwidDataV1 takes an hwid and returns the Sku and Variant in the form of
// HwidData proto. It will try the following in order:
// 1. Query from datastore. If under an hour old, return data.
// 2. If over an hour old or no data in datastore, attempt to query new data
//    from HWID server.
// 3. If HWID server data available, cache into datastore and return data.
// 4. If server fails, return expired datastore data if present. If not, return
//    nil and error.
func GetHwidDataV1(ctx context.Context, c hwid.ClientInterface, hwid string) (*ufspb.HwidData, error) {
	var hwidEnt *configuration.HwidDataEntity
	var hwidEntNew *configuration.HwidDataEntity

	hwidEnt, err := configuration.GetHwidData(ctx, hwid)
	if err != nil && !util.IsNotFoundError(err) {
		return nil, err
	}

	if hwidEnt == nil {
		hwidEntNew, err = cacheHwidData(ctx, c, hwid)
		if err != nil {
			return nil, err
		}
	} else {
		now := time.Now().UTC()
		lastUpdated := hwidEnt.Updated
		if now.After(lastUpdated.Add(cacheAge)) {
			hwidEntNew, err = cacheHwidData(ctx, c, hwid)
			if err != nil {
				logging.Warningf(ctx, "Error querying from HWID server. Falling back to existing HWID data: %s", err)
			}
		}
	}

	if hwidEntNew != nil {
		hwidEnt = hwidEntNew
	}

	data, err := configuration.ParseHwidDataV1(hwidEnt)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// cacheHwidData queries the hwid server with an hwid and stores the results
// into the UFS datastore.
func cacheHwidData(ctx context.Context, c hwid.ClientInterface, hwid string) (*configuration.HwidDataEntity, error) {
	newData, err := c.QueryHwid(ctx, hwid)
	if err != nil {
		return nil, err
	}

	resp, err := configuration.UpdateHwidData(ctx, newData, hwid)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
