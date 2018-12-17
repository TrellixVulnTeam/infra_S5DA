// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package config

import (
	"fmt"

	"go.chromium.org/luci/common/errors"

	"infra/tricium/api/v1"
)

// Validate validates the provided project config using the provided service config.
func Validate(sc *tricium.ServiceConfig, pc *tricium.ProjectConfig) (*tricium.ProjectConfig, error) {
	if sc.SwarmingServer == "" {
		return nil, errors.Reason("missing swarming server URL in service config").Err()
	}
	if sc.IsolateServer == "" {
		return nil, errors.Reason("missing isolate server URL in service config").Err()
	}
	if sc.BuildbucketServerHost == "" {
		return nil, errors.Reason("missing buildbucket server host in service config").Err()
	}
	if pc.SwarmingServiceAccount == "" {
		return nil, errors.Reason("missing swarming service account for project: %+v", pc).Err()
	}
	res := &tricium.ProjectConfig{
		Acls:       pc.Acls,
		Selections: pc.Selections,
	}
	functions := map[string]*tricium.Function{}
	for _, s := range pc.Selections {
		// Get merged function definition.
		if _, ok := functions[s.Function]; !ok {
			sca := tricium.LookupFunction(sc.Functions, s.Function)
			pca := tricium.LookupFunction(pc.Functions, s.Function)
			f, err := mergeFunctions(s.Function, sc, sca, pca)
			if err != nil {
				return nil, err
			}
			functions[s.Function] = f
			res.Functions = append(res.Functions, f)
		}
		if err := tricium.ValidateFunction(functions[s.Function], sc); err != nil {
			return nil, errors.Annotate(err, "function is not valid").Err()
		}
		if !tricium.SupportsPlatform(functions[s.Function], s.Platform) {
			return nil, errors.Reason("no support for platform %s by function %s", s.Platform, s.Function).Err()
		}
		for _, c := range s.Configs {
			if !tricium.SupportsConfig(functions[s.Function], c) {
				return nil, errors.Reason("no support for config %s by function %s", c.Name, s.Function).Err()
			}
		}
	}
	for _, v := range functions {
		res.Functions = append(res.Functions, v)
	}
	return res, nil
}

// mergeFunctions merges the provided service and project function configs.
//
// In merging, the following override rules are applied:
// - existence of project path_filters fully replaces any service path_filters.
// - project impl for platform fully replaces service impl for the same platform
//   project owner or component.
//
// Errors:
//  - change of data dependency in service config is not allowed
func mergeFunctions(function string, sc *tricium.ServiceConfig, sa, pa *tricium.Function) (*tricium.Function, error) {
	// TODO(qyearsley): Extract nil checks out of this function so that
	// it is more focused on merging.
	if sa == nil && pa == nil {
		return nil, errors.Reason("unknown function %s", function).Err()
	}
	res := &tricium.Function{Name: function}
	if sa != nil {
		if err := tricium.ValidateFunction(sa, sc); err != nil {
			return nil, errors.Reason("invalid service function config for %s: %v", function, err).Err()
		}
		if sa.GetNeeds() == tricium.Data_NONE || sa.GetProvides() == tricium.Data_NONE {
			return nil, errors.Reason("service function config must have data dependencies, function: %s", function).Err()
		}
		res.Type = sa.Type
		res.Needs = sa.Needs
		res.Provides = sa.Provides
		res.PathFilters = sa.PathFilters
		res.Owner = sa.Owner
		res.MonorailComponent = sa.MonorailComponent
		res.ConfigDefs = sa.ConfigDefs
		res.Impls = sa.Impls
	}
	if pa != nil {
		if sa != nil &&
			(sa.Type != pa.Type) {
			return nil, fmt.Errorf("cannot merge functions of different type, name: %s, service type: %s, project type: %s",
				pa.Name, sa.Type, pa.Type)
		}
		res.Type = pa.Type
		if sa != nil &&
			(pa.GetNeeds() != tricium.Data_NONE && pa.GetNeeds() != sa.GetNeeds() ||
				pa.GetProvides() != tricium.Data_NONE && pa.GetProvides() != sa.GetProvides()) {
			return nil, errors.Reason("change of service function data dependencies not allowed, function: %s", function).Err()
		}
		if sa == nil {
			if pa.GetNeeds() == tricium.Data_NONE || pa.GetProvides() == tricium.Data_NONE {
				return nil, errors.Reason("project function config is missing data dependencies, function: %s", function).Err()
			}
			res.Needs = pa.Needs
			res.Provides = pa.Provides
		}
		// Add service deps to project entry for validation check.
		// These deps are used to check validity of impls and when
		// there are service deps they are inherited by the project
		// entry.
		if sa != nil {
			pa.Needs = sa.Needs
			pa.Provides = sa.Provides
		}
		if err := tricium.ValidateFunction(pa, sc); err != nil {
			return nil, errors.Reason("invalid project function config for %s: %v", function, err).Err()
		}
		if pa.GetPathFilters() != nil {
			res.PathFilters = pa.PathFilters
		}
		if pa.GetOwner() != "" {
			res.Owner = pa.Owner
		}
		if pa.GetMonorailComponent() != "" {
			res.MonorailComponent = pa.MonorailComponent
		}
		if sa != nil {
			res.ConfigDefs = mergeConfigDefs(sa.ConfigDefs, pa.ConfigDefs)
			res.Impls = mergeImpls(sa.Impls, pa.Impls)
		} else {
			res.ConfigDefs = pa.ConfigDefs
			res.Impls = pa.Impls
		}
	}
	return res, nil
}

// mergeConfigDefs merges the service function config defs with the project
// function config defs.
//
// The project config defs can override service config defs with the same name.
func mergeConfigDefs(scd []*tricium.ConfigDef, pcd []*tricium.ConfigDef) []*tricium.ConfigDef {
	configs := map[string]*tricium.ConfigDef{}
	for _, cd := range scd {
		configs[cd.Name] = cd
	}
	for _, cd := range pcd {
		configs[cd.Name] = cd
	}
	res := []*tricium.ConfigDef{}
	for _, v := range configs {
		res = append(res, v)
	}
	return res
}

// mergeImpls merges the service function implementations with the project
// function implementations.
//
// All provided impl entries are assumed to be valid. The project
// implementations can override the service implementations for the same
// platform.
func mergeImpls(sci []*tricium.Impl, pci []*tricium.Impl) []*tricium.Impl {
	impls := map[tricium.Platform_Name]*tricium.Impl{}
	for _, i := range sci {
		impls[i.ProvidesForPlatform] = i
	}
	for _, i := range pci {
		impls[i.ProvidesForPlatform] = i
	}
	res := []*tricium.Impl{}
	for _, v := range impls {
		res = append(res, v)
	}
	return res
}
