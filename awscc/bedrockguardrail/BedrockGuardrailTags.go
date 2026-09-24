// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailTags struct {
	// Tag Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_guardrail#key BedrockGuardrail#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Tag Value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_guardrail#value BedrockGuardrail#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

