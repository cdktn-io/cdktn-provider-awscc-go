// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreconfigurationbundle


type BedrockagentcoreConfigurationBundleCreatedBy struct {
	// The Amazon Resource Name (ARN) of the source, if applicable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_configuration_bundle#arn BedrockagentcoreConfigurationBundle#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The name of the source (for example, user, optimization-job, or system).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_configuration_bundle#name BedrockagentcoreConfigurationBundle#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

