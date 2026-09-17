// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecuretunnelingtunnel


type IotsecuretunnelingTunnelDestinationConfig struct {
	// A list of service names that identify the target application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsecuretunneling_tunnel#services IotsecuretunnelingTunnel#services}
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
	// The name of the IoT thing to which you want to connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsecuretunneling_tunnel#thing_name IotsecuretunnelingTunnel#thing_name}
	ThingName *string `field:"optional" json:"thingName" yaml:"thingName"`
}

