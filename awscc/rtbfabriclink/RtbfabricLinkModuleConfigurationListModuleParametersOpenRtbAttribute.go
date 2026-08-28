// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabriclink


type RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttribute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_link#action RtbfabricLink#action}.
	Action *RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeAction `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_link#filter_configuration RtbfabricLink#filter_configuration}.
	FilterConfiguration interface{} `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_link#filter_type RtbfabricLink#filter_type}.
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_link#holdback_percentage RtbfabricLink#holdback_percentage}.
	HoldbackPercentage *float64 `field:"optional" json:"holdbackPercentage" yaml:"holdbackPercentage"`
}

