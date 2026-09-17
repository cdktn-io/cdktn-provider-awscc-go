// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationSelfServiceAiAgentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_agent#association_configurations WisdomAiAgent#association_configurations}.
	AssociationConfigurations interface{} `field:"optional" json:"associationConfigurations" yaml:"associationConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_agent#self_service_ai_guardrail_id WisdomAiAgent#self_service_ai_guardrail_id}.
	SelfServiceAiGuardrailId *string `field:"optional" json:"selfServiceAiGuardrailId" yaml:"selfServiceAiGuardrailId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_agent#self_service_answer_generation_ai_prompt_id WisdomAiAgent#self_service_answer_generation_ai_prompt_id}.
	SelfServiceAnswerGenerationAiPromptId *string `field:"optional" json:"selfServiceAnswerGenerationAiPromptId" yaml:"selfServiceAnswerGenerationAiPromptId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_ai_agent#self_service_pre_processing_ai_prompt_id WisdomAiAgent#self_service_pre_processing_ai_prompt_id}.
	SelfServicePreProcessingAiPromptId *string `field:"optional" json:"selfServicePreProcessingAiPromptId" yaml:"selfServicePreProcessingAiPromptId"`
}

