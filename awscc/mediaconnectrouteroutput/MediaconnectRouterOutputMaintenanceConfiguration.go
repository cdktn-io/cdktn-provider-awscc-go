// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputMaintenanceConfiguration struct {
	// Configuration settings for default maintenance scheduling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_router_output#default MediaconnectRouterOutput#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Configuration for preferred day and time maintenance settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediaconnect_router_output#preferred_day_time MediaconnectRouterOutput#preferred_day_time}
	PreferredDayTime *MediaconnectRouterOutputMaintenanceConfigurationPreferredDayTime `field:"optional" json:"preferredDayTime" yaml:"preferredDayTime"`
}

