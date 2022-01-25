// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package frontend

import (
	"context"
	"fmt"
	"math"
	"strings"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"

	fleet "infra/appengine/crosskylabadmin/api/fleet/v1"
	"infra/appengine/crosskylabadmin/internal/app/clients"
	"infra/appengine/crosskylabadmin/internal/app/config"
	"infra/appengine/crosskylabadmin/internal/app/frontend/internal/worker"
	"infra/libs/skylab/common/heuristics"
)

// Paris represents a decision to use the paris stack for this request.
const paris = "paris"

// Legacy represents a decision to use the legacy stack for this request.
const legacy = "legacy"

// Reason is a rationale for why we made the decision that we made.
type reason int

const (
	parisNotEnabled reason = iota
	allLabstationsAreOptedIn
	noPools
	wrongPool
	scoreExceedsThreshold
	scoreTooLow
	thresholdZero
	malformedPolicy
)

// ReasonMessageMap maps each reason to a readable description.
var reasonMessageMap = map[reason]string{
	parisNotEnabled:          "PARIS is not enabled",
	allLabstationsAreOptedIn: "All Labstations are opted in",
	noPools:                  "Labstation has no pools, possibly due to error calling UFS",
	wrongPool:                "Labstation has a pool not matching opted-in pools",
	scoreExceedsThreshold:    "Random score associated with task exceeds threshold",
	scoreTooLow:              "Random score associated with task does NOT exceed threshold",
	thresholdZero:            "Route labstation repair task: a threshold of zero implies that optinAllLabstations should be set, but optinAllLabstations is not set",
	malformedPolicy:          "Unrecognized policy",
}

// UFSErrorPolicy controls how UFS errors are handled.
type ufsErrorPolicy string

// UFS error policy constants.
// Error policy constants are defined in go/src/infra/appengine/crosskylabadmin/app/config/config.proto.
//
// Strict   -- fail on UFS error even if we don't need the result
// Fallback -- if we encounter a UFS error, fall back to the legacy path.
// Lax      -- if we do not need the UFS response to make a decision, do not fail the request.
const (
	// The strict policy causes all UFS error requests to be treated as fatal and causes the request to fail.
	ufsErrorPolicyStrict   ufsErrorPolicy = "strict"
	ufsErrorPolicyFallback                = "fallback"
	ufsErrorPolicyLax                     = "lax"
)

// NormalizeError policy normalizes a string into the canonical name for a policy.
func normalizeErrorPolicy(policy string) (ufsErrorPolicy, error) {
	policy = strings.ToLower(policy)
	switch policy {
	case "", "default", "fallback":
		return ufsErrorPolicyFallback, nil
	case "strict":
		return ufsErrorPolicyStrict, nil
	case "lax":
		return ufsErrorPolicyLax, nil
	}
	return "", fmt.Errorf("unrecognized policy: %q", policy)
}

// RouteRepairTask routes a repair task for a given bot.
//
// The possible return values are:
// - "legacy"  (for legacy, which is the default)
// -       ""  (indicates an error, should be treated as equivalent to "legacy" by callers)
// -  "paris"  (for PARIS, which is new)
//
// RouteRepairTask takes as an argument randFloat (which is a float64 in the closed interval [0, 1]).
// This argument is, by design, all the entropy that randFloat will need. Taking this as an argument allows
// RouteRepairTask itself to be deterministic because the caller is responsible for generating the random
// value.
func RouteRepairTask(ctx context.Context, botID string, expectedState string, pools []string, randFloat float64) (string, error) {
	if !(0.0 <= randFloat && randFloat <= 1.0) {
		return "", fmt.Errorf("Route repair task: randfloat %f is not in [0, 1]", randFloat)
	}
	if heuristics.LooksLikeLabstation(botID) {
		out, r := routeLabstationRepairTask(ctx, config.Get(ctx).GetParis().GetLabstationRepair(), pools, randFloat)
		reason, ok := reasonMessageMap[r]
		if !ok {
			logging.Infof(ctx, "Unrecognized reason %d", int64(r))
		}
		logging.Infof(ctx, "Sending labstation repair to %q because %q", out, reason)
		return out, nil
	}
	logging.Infof(ctx, "Non-labstations always use legacy flow")
	return "", nil
}

// CreateRepairTask kicks off a repair job.
//
// This function will either schedule a legacy repair task or a PARIS repair task.
// Note that the ufs client can be nil.
func CreateRepairTask(ctx context.Context, botID string, expectedState string, pools []string, randFloat float64) (string, error) {
	logging.Infof(ctx, "Creating repair task for %q expected state %q with random input %f", botID, expectedState, randFloat)
	// If we encounter an error picking paris or legacy, do the safe thing and use legacy.
	taskType, err := RouteRepairTask(ctx, botID, expectedState, pools, randFloat)
	if err != nil {
		logging.Infof(ctx, "Create repair task: falling back to legacy repair by default: %s", err)
	}
	switch taskType {
	case "paris":
		return createBuildbucketRepairTask(ctx, botID, expectedState)
	default:
		return createLegacyRepairTask(ctx, botID, expectedState)
	}
}

// RouteLabstationRepairTask takes a repair task for a labstation and routes it.
// TODO(gregorynisbet): Generalize this to non-labstation tasks.
func routeLabstationRepairTask(ctx context.Context, r *config.RolloutConfig, pools []string, randFloat float64) (string, reason) {
	// TODO(gregorynisbet): Log instead of silently falling back to the default error handling policy.
	ufsErrorPolicy, err := normalizeErrorPolicy(r.GetUfsErrorPolicy())
	if err != nil {
		logging.Infof(ctx, "error while routing labstation repair task: %s", err)
	}
	// Check that the feature is enabled at all.
	if !r.GetEnable() {
		return legacy, parisNotEnabled
	}
	// Check for malformed input data that would cause us to be unable to make a decision.
	if len(pools) == 0 {
		switch ufsErrorPolicy {
		case ufsErrorPolicyLax:
			// Do nothing.
		case ufsErrorPolicyStrict:
			// TODO(gregorynisbet): Make strict error handling propagate the failure back up.
			return legacy, noPools
		case ufsErrorPolicyFallback:
			return legacy, noPools
		default:
			return legacy, malformedPolicy
		}
	}
	threshold := r.GetRolloutPermille()
	myValue := math.Round(1000.0 * randFloat)
	if r.GetOptinAllDuts() {
		if myValue >= float64(threshold) {
			return paris, scoreExceedsThreshold
		}
		return legacy, scoreTooLow
	}
	if threshold == 0 {
		return legacy, thresholdZero
	}
	if !r.GetOptinAllDuts() && len(r.GetOptinDutPool()) > 0 && isDisjoint(pools, r.GetOptinDutPool()) {
		return legacy, wrongPool
	}
	if myValue >= float64(threshold) {
		return paris, scoreExceedsThreshold
	}
	return legacy, scoreTooLow
}

// CreateBuildbucketRepairTask creates a new repair task for a labstation. Not yet implemented.
func createBuildbucketRepairTask(ctx context.Context, botID string, expectedState string) (string, error) {
	logging.Infof(ctx, "Using new repair flow for bot %q", botID)
	return "", fmt.Errorf("not yet implemented")
}

// CreateLegacyRepairTask creates a legacy repair task for a labstation.
func createLegacyRepairTask(ctx context.Context, botID string, expectedState string) (string, error) {
	logging.Infof(ctx, "Using legacy repair flow for bot %q", botID)
	at := worker.AdminTaskForType(ctx, fleet.TaskType_Repair)
	sc, err := clients.NewSwarmingClient(ctx, config.Get(ctx).Swarming.Host)
	if err != nil {
		return "", errors.Annotate(err, "failed to obtain swarming client").Err()
	}
	cfg := config.Get(ctx)
	taskURL, err := runTaskByBotID(ctx, at, sc, botID, expectedState, cfg.Tasker.BackgroundTaskExpirationSecs, cfg.Tasker.BackgroundTaskExecutionTimeoutSecs)
	if err != nil {
		return "", errors.Annotate(err, "fail to create repair task for %s", botID).Err()
	}
	return taskURL, nil
}

// IsRecoveryEnabledForLabstation returns whether recovery is enabled for labstations.
// TODO(gregorynisbet): Expand this to take into account other relevant factors like the pools that the labstation is in.
func isRecoveryEnabledForLabstation(ctx context.Context) bool {
	enabled := true
	enabled = enabled && config.Get(ctx).GetParis().GetLabstationRepair().GetEnable()
	return enabled
}

// IsDisjoint returns true if and only if two sequences have no elements in common.
func isDisjoint(a []string, b []string) bool {
	bMap := make(map[string]bool, len(b))
	for _, item := range b {
		bMap[item] = true
	}
	for _, item := range a {
		if bMap[item] {
			return false
		}
	}
	return true
}
