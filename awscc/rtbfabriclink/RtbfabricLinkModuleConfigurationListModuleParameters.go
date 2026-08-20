// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabriclink


type RtbfabricLinkModuleConfigurationListModuleParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rtbfabric_link#no_bid RtbfabricLink#no_bid}.
	NoBid *RtbfabricLinkModuleConfigurationListModuleParametersNoBid `field:"optional" json:"noBid" yaml:"noBid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rtbfabric_link#open_rtb_attribute RtbfabricLink#open_rtb_attribute}.
	OpenRtbAttribute *RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttribute `field:"optional" json:"openRtbAttribute" yaml:"openRtbAttribute"`
}

