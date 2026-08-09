// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailContentPolicyConfigFiltersConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_guardrail#input_action BedrockGuardrail#input_action}.
	InputAction *string `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_guardrail#input_enabled BedrockGuardrail#input_enabled}.
	InputEnabled interface{} `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// List of modalities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_guardrail#input_modalities BedrockGuardrail#input_modalities}
	InputModalities *[]*string `field:"optional" json:"inputModalities" yaml:"inputModalities"`
	// Strength for filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_guardrail#input_strength BedrockGuardrail#input_strength}
	InputStrength *string `field:"optional" json:"inputStrength" yaml:"inputStrength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_guardrail#output_action BedrockGuardrail#output_action}.
	OutputAction *string `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_guardrail#output_enabled BedrockGuardrail#output_enabled}.
	OutputEnabled interface{} `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
	// List of modalities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_guardrail#output_modalities BedrockGuardrail#output_modalities}
	OutputModalities *[]*string `field:"optional" json:"outputModalities" yaml:"outputModalities"`
	// Strength for filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_guardrail#output_strength BedrockGuardrail#output_strength}
	OutputStrength *string `field:"optional" json:"outputStrength" yaml:"outputStrength"`
	// Type of filter in content policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_guardrail#type BedrockGuardrail#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

