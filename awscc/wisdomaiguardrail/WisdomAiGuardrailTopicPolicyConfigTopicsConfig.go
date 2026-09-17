// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiguardrail


type WisdomAiGuardrailTopicPolicyConfigTopicsConfig struct {
	// Definition of topic in topic policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#definition WisdomAiGuardrail#definition}
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// List of text examples.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#examples WisdomAiGuardrail#examples}
	Examples *[]*string `field:"optional" json:"examples" yaml:"examples"`
	// Name of topic in topic policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#name WisdomAiGuardrail#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Type of topic in a policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#type WisdomAiGuardrail#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

