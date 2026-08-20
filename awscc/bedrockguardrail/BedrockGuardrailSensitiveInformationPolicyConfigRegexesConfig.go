// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailSensitiveInformationPolicyConfigRegexesConfig struct {
	// Options for sensitive information action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_guardrail#action BedrockGuardrail#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The regex description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_guardrail#description BedrockGuardrail#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Options for sensitive information action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_guardrail#input_action BedrockGuardrail#input_action}
	InputAction *string `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_guardrail#input_enabled BedrockGuardrail#input_enabled}.
	InputEnabled interface{} `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// The regex name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_guardrail#name BedrockGuardrail#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Options for sensitive information action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_guardrail#output_action BedrockGuardrail#output_action}
	OutputAction *string `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_guardrail#output_enabled BedrockGuardrail#output_enabled}.
	OutputEnabled interface{} `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
	// The regex pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_guardrail#pattern BedrockGuardrail#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

