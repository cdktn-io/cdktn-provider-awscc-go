// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabriclink


type RtbfabricLinkModuleConfigurationListModuleParametersNoBid struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rtbfabric_link#pass_through_percentage RtbfabricLink#pass_through_percentage}.
	PassThroughPercentage *float64 `field:"optional" json:"passThroughPercentage" yaml:"passThroughPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rtbfabric_link#reason RtbfabricLink#reason}.
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rtbfabric_link#reason_code RtbfabricLink#reason_code}.
	ReasonCode *float64 `field:"optional" json:"reasonCode" yaml:"reasonCode"`
}

