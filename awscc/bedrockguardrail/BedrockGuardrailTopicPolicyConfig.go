// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailTopicPolicyConfig struct {
	// List of topic configs in topic policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_guardrail#topics_config BedrockGuardrail#topics_config}
	TopicsConfig interface{} `field:"optional" json:"topicsConfig" yaml:"topicsConfig"`
	// Guardrail tier config for topic policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_guardrail#topics_tier_config BedrockGuardrail#topics_tier_config}
	TopicsTierConfig *BedrockGuardrailTopicPolicyConfigTopicsTierConfig `field:"optional" json:"topicsTierConfig" yaml:"topicsTierConfig"`
}

