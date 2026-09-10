// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockGuardrailConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Messaging for when violations are detected in text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#blocked_input_messaging BedrockGuardrail#blocked_input_messaging}
	BlockedInputMessaging *string `field:"required" json:"blockedInputMessaging" yaml:"blockedInputMessaging"`
	// Messaging for when violations are detected in text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#blocked_outputs_messaging BedrockGuardrail#blocked_outputs_messaging}
	BlockedOutputsMessaging *string `field:"required" json:"blockedOutputsMessaging" yaml:"blockedOutputsMessaging"`
	// Name of the guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#name BedrockGuardrail#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional configuration for integrating Automated Reasoning policies with the guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#automated_reasoning_policy_config BedrockGuardrail#automated_reasoning_policy_config}
	AutomatedReasoningPolicyConfig *BedrockGuardrailAutomatedReasoningPolicyConfig `field:"optional" json:"automatedReasoningPolicyConfig" yaml:"automatedReasoningPolicyConfig"`
	// Content policy config for a guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#content_policy_config BedrockGuardrail#content_policy_config}
	ContentPolicyConfig *BedrockGuardrailContentPolicyConfig `field:"optional" json:"contentPolicyConfig" yaml:"contentPolicyConfig"`
	// Contextual grounding policy config for a guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#contextual_grounding_policy_config BedrockGuardrail#contextual_grounding_policy_config}
	ContextualGroundingPolicyConfig *BedrockGuardrailContextualGroundingPolicyConfig `field:"optional" json:"contextualGroundingPolicyConfig" yaml:"contextualGroundingPolicyConfig"`
	// The system-defined guardrail profile that you?re using with your guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#cross_region_config BedrockGuardrail#cross_region_config}
	CrossRegionConfig *BedrockGuardrailCrossRegionConfig `field:"optional" json:"crossRegionConfig" yaml:"crossRegionConfig"`
	// Description of the guardrail or its version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#description BedrockGuardrail#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key with which the guardrail was encrypted at rest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#kms_key_arn BedrockGuardrail#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Sensitive information policy config for a guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#sensitive_information_policy_config BedrockGuardrail#sensitive_information_policy_config}
	SensitiveInformationPolicyConfig *BedrockGuardrailSensitiveInformationPolicyConfig `field:"optional" json:"sensitiveInformationPolicyConfig" yaml:"sensitiveInformationPolicyConfig"`
	// List of Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#tags BedrockGuardrail#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Topic policy config for a guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#topic_policy_config BedrockGuardrail#topic_policy_config}
	TopicPolicyConfig *BedrockGuardrailTopicPolicyConfig `field:"optional" json:"topicPolicyConfig" yaml:"topicPolicyConfig"`
	// Word policy config for a guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#word_policy_config BedrockGuardrail#word_policy_config}
	WordPolicyConfig *BedrockGuardrailWordPolicyConfig `field:"optional" json:"wordPolicyConfig" yaml:"wordPolicyConfig"`
}

