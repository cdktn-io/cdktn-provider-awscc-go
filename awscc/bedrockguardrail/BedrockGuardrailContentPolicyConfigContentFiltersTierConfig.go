// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailContentPolicyConfigContentFiltersTierConfig struct {
	// Tier name for tier configuration in content filters policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_guardrail#tier_name BedrockGuardrail#tier_name}
	TierName *string `field:"optional" json:"tierName" yaml:"tierName"`
}

