// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationCaseSummarizationAiAgentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#case_summarization_ai_guardrail_id WisdomAiAgent#case_summarization_ai_guardrail_id}.
	CaseSummarizationAiGuardrailId *string `field:"optional" json:"caseSummarizationAiGuardrailId" yaml:"caseSummarizationAiGuardrailId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#case_summarization_ai_prompt_id WisdomAiAgent#case_summarization_ai_prompt_id}.
	CaseSummarizationAiPromptId *string `field:"optional" json:"caseSummarizationAiPromptId" yaml:"caseSummarizationAiPromptId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#locale WisdomAiAgent#locale}.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
}

