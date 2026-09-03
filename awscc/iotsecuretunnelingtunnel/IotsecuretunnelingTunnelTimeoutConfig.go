// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecuretunnelingtunnel


type IotsecuretunnelingTunnelTimeoutConfig struct {
	// The maximum amount of time (in minutes) a tunnel can remain open.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsecuretunneling_tunnel#max_lifetime_timeout_minutes IotsecuretunnelingTunnel#max_lifetime_timeout_minutes}
	MaxLifetimeTimeoutMinutes *float64 `field:"optional" json:"maxLifetimeTimeoutMinutes" yaml:"maxLifetimeTimeoutMinutes"`
}

