// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package outpostssite


type OutpostsSiteRackPhysicalProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/outposts_site#fiber_optic_cable_type OutpostsSite#fiber_optic_cable_type}.
	FiberOpticCableType *string `field:"optional" json:"fiberOpticCableType" yaml:"fiberOpticCableType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/outposts_site#maximum_supported_weight_lbs OutpostsSite#maximum_supported_weight_lbs}.
	MaximumSupportedWeightLbs *string `field:"optional" json:"maximumSupportedWeightLbs" yaml:"maximumSupportedWeightLbs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/outposts_site#optical_standard OutpostsSite#optical_standard}.
	OpticalStandard *string `field:"optional" json:"opticalStandard" yaml:"opticalStandard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/outposts_site#power_connector OutpostsSite#power_connector}.
	PowerConnector *string `field:"optional" json:"powerConnector" yaml:"powerConnector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/outposts_site#power_draw_kva OutpostsSite#power_draw_kva}.
	PowerDrawKva *string `field:"optional" json:"powerDrawKva" yaml:"powerDrawKva"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/outposts_site#power_feed_drop OutpostsSite#power_feed_drop}.
	PowerFeedDrop *string `field:"optional" json:"powerFeedDrop" yaml:"powerFeedDrop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/outposts_site#power_phase OutpostsSite#power_phase}.
	PowerPhase *string `field:"optional" json:"powerPhase" yaml:"powerPhase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/outposts_site#uplink_count OutpostsSite#uplink_count}.
	UplinkCount *string `field:"optional" json:"uplinkCount" yaml:"uplinkCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/outposts_site#uplink_gbps OutpostsSite#uplink_gbps}.
	UplinkGbps *string `field:"optional" json:"uplinkGbps" yaml:"uplinkGbps"`
}

