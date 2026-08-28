// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockautomatedreasoningpolicy


type BedrockAutomatedReasoningPolicyPolicyDefinitionTypes struct {
	// A natural language description of this type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_automated_reasoning_policy#description BedrockAutomatedReasoningPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for this type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_automated_reasoning_policy#name BedrockAutomatedReasoningPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of valid values for this type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_automated_reasoning_policy#values BedrockAutomatedReasoningPolicy#values}
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

