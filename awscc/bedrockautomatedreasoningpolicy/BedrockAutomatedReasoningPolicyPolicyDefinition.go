// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockautomatedreasoningpolicy


type BedrockAutomatedReasoningPolicyPolicyDefinition struct {
	// The rules definition block of an AutomatedReasoningPolicyDefinition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_automated_reasoning_policy#rules BedrockAutomatedReasoningPolicy#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// The types definition block of an AutomatedReasoningPolicyDefinition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_automated_reasoning_policy#types BedrockAutomatedReasoningPolicy#types}
	Types interface{} `field:"optional" json:"types" yaml:"types"`
	// The variables definition block of an AutomatedReasoningPolicyDefinition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_automated_reasoning_policy#variables BedrockAutomatedReasoningPolicy#variables}
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
	// The policy format version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_automated_reasoning_policy#version BedrockAutomatedReasoningPolicy#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

