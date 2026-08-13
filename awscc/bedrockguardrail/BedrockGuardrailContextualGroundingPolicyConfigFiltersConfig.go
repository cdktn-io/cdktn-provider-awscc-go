// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailContextualGroundingPolicyConfigFiltersConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_guardrail#action BedrockGuardrail#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_guardrail#enabled BedrockGuardrail#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The threshold for this filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_guardrail#threshold BedrockGuardrail#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// Type of contextual grounding filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_guardrail#type BedrockGuardrail#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

