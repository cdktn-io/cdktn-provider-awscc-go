// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputMaintenanceConfiguration struct {
	// Configuration settings for default maintenance scheduling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_router_input#default MediaconnectRouterInput#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Configuration for preferred day and time maintenance settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_router_input#preferred_day_time MediaconnectRouterInput#preferred_day_time}
	PreferredDayTime *MediaconnectRouterInputMaintenanceConfigurationPreferredDayTime `field:"optional" json:"preferredDayTime" yaml:"preferredDayTime"`
}

