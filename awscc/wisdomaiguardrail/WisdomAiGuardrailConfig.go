// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiguardrail

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomAiGuardrailConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#assistant_id WisdomAiGuardrail#assistant_id}.
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// Messaging for when violations are detected in text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#blocked_input_messaging WisdomAiGuardrail#blocked_input_messaging}
	BlockedInputMessaging *string `field:"required" json:"blockedInputMessaging" yaml:"blockedInputMessaging"`
	// Messaging for when violations are detected in text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#blocked_outputs_messaging WisdomAiGuardrail#blocked_outputs_messaging}
	BlockedOutputsMessaging *string `field:"required" json:"blockedOutputsMessaging" yaml:"blockedOutputsMessaging"`
	// Content policy config for a guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#content_policy_config WisdomAiGuardrail#content_policy_config}
	ContentPolicyConfig *WisdomAiGuardrailContentPolicyConfig `field:"optional" json:"contentPolicyConfig" yaml:"contentPolicyConfig"`
	// Contextual grounding policy config for a guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#contextual_grounding_policy_config WisdomAiGuardrail#contextual_grounding_policy_config}
	ContextualGroundingPolicyConfig *WisdomAiGuardrailContextualGroundingPolicyConfig `field:"optional" json:"contextualGroundingPolicyConfig" yaml:"contextualGroundingPolicyConfig"`
	// Description of the guardrail or its version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#description WisdomAiGuardrail#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#name WisdomAiGuardrail#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Sensitive information policy config for a guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#sensitive_information_policy_config WisdomAiGuardrail#sensitive_information_policy_config}
	SensitiveInformationPolicyConfig *WisdomAiGuardrailSensitiveInformationPolicyConfig `field:"optional" json:"sensitiveInformationPolicyConfig" yaml:"sensitiveInformationPolicyConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#tags WisdomAiGuardrail#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Topic policy config for a guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#topic_policy_config WisdomAiGuardrail#topic_policy_config}
	TopicPolicyConfig *WisdomAiGuardrailTopicPolicyConfig `field:"optional" json:"topicPolicyConfig" yaml:"topicPolicyConfig"`
	// Word policy config for a guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_guardrail#word_policy_config WisdomAiGuardrail#word_policy_config}
	WordPolicyConfig *WisdomAiGuardrailWordPolicyConfig `field:"optional" json:"wordPolicyConfig" yaml:"wordPolicyConfig"`
}

