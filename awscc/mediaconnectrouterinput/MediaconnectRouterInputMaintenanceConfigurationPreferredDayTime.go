// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputMaintenanceConfigurationPreferredDayTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_router_input#day MediaconnectRouterInput#day}.
	Day *string `field:"optional" json:"day" yaml:"day"`
	// The preferred time for maintenance operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_router_input#time MediaconnectRouterInput#time}
	Time *string `field:"optional" json:"time" yaml:"time"`
}

