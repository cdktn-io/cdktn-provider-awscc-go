// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailContextualGroundingPolicyConfig struct {
	// List of contextual grounding filter configs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_guardrail#filters_config BedrockGuardrail#filters_config}
	FiltersConfig interface{} `field:"optional" json:"filtersConfig" yaml:"filtersConfig"`
}

