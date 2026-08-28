// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockautomatedreasoningpolicy


type BedrockAutomatedReasoningPolicyPolicyDefinitionTypesValues struct {
	// A natural language description of the type's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_automated_reasoning_policy#description BedrockAutomatedReasoningPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The value of the type value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_automated_reasoning_policy#value BedrockAutomatedReasoningPolicy#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

