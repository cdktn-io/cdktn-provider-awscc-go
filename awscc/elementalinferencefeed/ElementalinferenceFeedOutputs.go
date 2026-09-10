// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elementalinferencefeed


type ElementalinferenceFeedOutputs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elementalinference_feed#name ElementalinferenceFeed#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elementalinference_feed#output_config ElementalinferenceFeed#output_config}.
	OutputConfig *ElementalinferenceFeedOutputsOutputConfig `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elementalinference_feed#status ElementalinferenceFeed#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elementalinference_feed#description ElementalinferenceFeed#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

