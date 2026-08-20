// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiguardrail


type WisdomAiGuardrailSensitiveInformationPolicyConfigRegexesConfig struct {
	// Options for sensitive information action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_guardrail#action WisdomAiGuardrail#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The regex description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_guardrail#description WisdomAiGuardrail#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The regex name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_guardrail#name WisdomAiGuardrail#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The regex pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_guardrail#pattern WisdomAiGuardrail#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

