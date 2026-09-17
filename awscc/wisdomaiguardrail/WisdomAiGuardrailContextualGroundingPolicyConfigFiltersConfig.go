// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiguardrail


type WisdomAiGuardrailContextualGroundingPolicyConfigFiltersConfig struct {
	// The threshold for this filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#threshold WisdomAiGuardrail#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// Type of contextual grounding filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#type WisdomAiGuardrail#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

