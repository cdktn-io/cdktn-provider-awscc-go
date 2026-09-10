// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockautomatedreasoningpolicy


type BedrockAutomatedReasoningPolicyPolicyDefinitionVariables struct {
	// A natural language description of this variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_automated_reasoning_policy#description BedrockAutomatedReasoningPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name from this variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_automated_reasoning_policy#name BedrockAutomatedReasoningPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A type for this variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_automated_reasoning_policy#type BedrockAutomatedReasoningPolicy#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

