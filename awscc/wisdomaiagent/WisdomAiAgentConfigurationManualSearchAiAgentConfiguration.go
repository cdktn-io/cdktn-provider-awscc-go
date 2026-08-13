// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationManualSearchAiAgentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/wisdom_ai_agent#answer_generation_ai_guardrail_id WisdomAiAgent#answer_generation_ai_guardrail_id}.
	AnswerGenerationAiGuardrailId *string `field:"optional" json:"answerGenerationAiGuardrailId" yaml:"answerGenerationAiGuardrailId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/wisdom_ai_agent#answer_generation_ai_prompt_id WisdomAiAgent#answer_generation_ai_prompt_id}.
	AnswerGenerationAiPromptId *string `field:"optional" json:"answerGenerationAiPromptId" yaml:"answerGenerationAiPromptId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/wisdom_ai_agent#association_configurations WisdomAiAgent#association_configurations}.
	AssociationConfigurations interface{} `field:"optional" json:"associationConfigurations" yaml:"associationConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/wisdom_ai_agent#locale WisdomAiAgent#locale}.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
}

