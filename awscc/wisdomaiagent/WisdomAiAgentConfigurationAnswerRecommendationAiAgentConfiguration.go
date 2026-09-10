// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#answer_generation_ai_guardrail_id WisdomAiAgent#answer_generation_ai_guardrail_id}.
	AnswerGenerationAiGuardrailId *string `field:"optional" json:"answerGenerationAiGuardrailId" yaml:"answerGenerationAiGuardrailId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#answer_generation_ai_prompt_id WisdomAiAgent#answer_generation_ai_prompt_id}.
	AnswerGenerationAiPromptId *string `field:"optional" json:"answerGenerationAiPromptId" yaml:"answerGenerationAiPromptId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#association_configurations WisdomAiAgent#association_configurations}.
	AssociationConfigurations interface{} `field:"optional" json:"associationConfigurations" yaml:"associationConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#intent_labeling_generation_ai_prompt_id WisdomAiAgent#intent_labeling_generation_ai_prompt_id}.
	IntentLabelingGenerationAiPromptId *string `field:"optional" json:"intentLabelingGenerationAiPromptId" yaml:"intentLabelingGenerationAiPromptId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#locale WisdomAiAgent#locale}.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#query_reformulation_ai_prompt_id WisdomAiAgent#query_reformulation_ai_prompt_id}.
	QueryReformulationAiPromptId *string `field:"optional" json:"queryReformulationAiPromptId" yaml:"queryReformulationAiPromptId"`
}

