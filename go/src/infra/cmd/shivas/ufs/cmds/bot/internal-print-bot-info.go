// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cmds

import (
	"encoding/json"
	"fmt"

	"github.com/maruel/subcommands"
	"go.chromium.org/luci/auth/client/authcli"
	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/grpc/prpc"

	"infra/cmd/shivas/site"
	"infra/cmd/shivas/utils"
	"infra/cmdsupport/cmdlib"
	"infra/cros/dutstate"
	"infra/libs/skylab/inventory"
	"infra/libs/skylab/inventory/swarming"
	ufsAPI "infra/unifiedfleet/api/v1/rpc"
	ufsUtil "infra/unifiedfleet/app/util"
	ufsOSUtil "infra/unifiedfleet/app/util/osutil"
)

// PrintBotInfo subcommand: Print Swarming dimensions for a DUT.
var PrintBotInfo = &subcommands.Command{
	UsageLine: "internal-print-bot-info DUT hostname/Asset tag",
	ShortDesc: "print Swarming bot info for a DUT",
	LongDesc: `Print Swarming bot info for a DUT.

For internal use only.`,
	CommandRun: func() subcommands.CommandRun {
		c := &printBotInfoRun{}
		c.authFlags.Register(&c.Flags, site.DefaultAuthOptions)
		c.envFlags.Register(&c.Flags)
		c.commonFlags.Register(&c.Flags)

		c.Flags.BoolVar(&c.byHostname, "by-hostname", false, "Lookup by hostname instead of ID/Asset tag.")
		return c
	},
}

type printBotInfoRun struct {
	subcommands.CommandRunBase
	authFlags   authcli.Flags
	envFlags    site.EnvFlags
	commonFlags site.CommonFlags

	byHostname bool
}

func (c *printBotInfoRun) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	if err := c.innerRun(a, args, env); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	return 0
}

func (c *printBotInfoRun) innerRun(a subcommands.Application, args []string, env subcommands.Env) error {
	if len(args) != 1 {
		return cmdlib.NewUsageError(c.Flags, "exactly one DUT hostname must be provided")
	}
	ctx := cli.GetContext(a, c, env)
	ctx = utils.SetupContext(ctx, ufsUtil.OSNamespace)
	hc, err := cmdlib.NewHTTPClient(ctx, &c.authFlags)
	if err != nil {
		return err
	}
	e := c.envFlags.Env()
	if c.commonFlags.Verbose() {
		fmt.Printf("Using UnifiedFleet service %s\n", e.UnifiedFleetService)
	}
	ufsClient := ufsAPI.NewFleetPRPCClient(&prpc.Client{
		C:       hc,
		Host:    e.UnifiedFleetService,
		Options: site.DefaultPRPCOptions,
	})
	req := &ufsAPI.GetChromeOSDeviceDataRequest{}
	if c.byHostname {
		req.Hostname = args[0]
	} else {
		req.ChromeosDeviceId = args[0]
	}
	data, err := ufsClient.GetChromeOSDeviceData(ctx, req)
	if err != nil {
		return err
	}
	d, err := ufsOSUtil.AdaptToV1DutSpec(data)
	if err != nil {
		return err
	}
	dutStateInfo := dutstate.Read(ctx, ufsClient, data.GetLabConfig().GetName())
	stderr := a.GetErr()
	r := func(e error) { fmt.Fprintf(stderr, "sanitize dimensions: %s\n", err) }
	bi := botInfoForDUT(d, dutStateInfo, r)
	enc, err := json.Marshal(bi)
	if err != nil {
		return err
	}
	a.GetOut().Write(enc)
	return nil
}

type botInfo struct {
	Dimensions swarming.Dimensions
	State      botState
}

type botState map[string][]string

func botInfoForDUT(d *inventory.DeviceUnderTest, ds dutstate.Info, r swarming.ReportFunc) botInfo {
	return botInfo{
		Dimensions: botDimensionsForDUT(d, ds, r),
		State:      botStateForDUT(d),
	}
}

func botStateForDUT(d *inventory.DeviceUnderTest) botState {
	s := make(botState)
	for _, kv := range d.GetCommon().GetAttributes() {
		k, v := kv.GetKey(), kv.GetValue()
		s[k] = append(s[k], v)
	}
	s["storage_state"] = []string{d.GetCommon().GetLabels().GetPeripherals().GetStorageState().String()[len("HARDWARE_"):]}
	s["servo_usb_state"] = []string{d.GetCommon().GetLabels().GetPeripherals().GetServoUsbState().String()[len("HARDWARE_"):]}
	s["rpm_state"] = []string{d.GetCommon().GetLabels().GetPeripherals().GetRpmState().String()}
	return s
}

func botDimensionsForDUT(d *inventory.DeviceUnderTest, ds dutstate.Info, r swarming.ReportFunc) swarming.Dimensions {
	c := d.GetCommon()
	dims := swarming.Convert(c.GetLabels())
	dims["dut_id"] = []string{c.GetId()}
	dims["dut_name"] = []string{c.GetHostname()}
	if v := c.GetHwid(); v != "" {
		dims["hwid"] = []string{v}
	}
	if v := c.GetSerialNumber(); v != "" {
		dims["serial_number"] = []string{v}
	}
	if v := c.GetLocation(); v != nil {
		dims["location"] = []string{formatLocation(v)}
	}
	dims["dut_state"] = []string{string(ds.State)}
	swarming.Sanitize(dims, r)
	return dims
}

func formatLocation(loc *inventory.Location) string {
	return fmt.Sprintf("%s-row%d-rack%d-host%d",
		loc.GetLab().GetName(),
		loc.GetRow(),
		loc.GetRack(),
		loc.GetHost(),
	)
}
