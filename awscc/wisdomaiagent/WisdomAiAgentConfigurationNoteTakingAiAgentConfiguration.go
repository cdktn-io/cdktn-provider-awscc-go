// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationNoteTakingAiAgentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_ai_agent#locale WisdomAiAgent#locale}.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_ai_agent#note_taking_ai_guardrail_id WisdomAiAgent#note_taking_ai_guardrail_id}.
	NoteTakingAiGuardrailId *string `field:"optional" json:"noteTakingAiGuardrailId" yaml:"noteTakingAiGuardrailId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_ai_agent#note_taking_ai_prompt_id WisdomAiAgent#note_taking_ai_prompt_id}.
	NoteTakingAiPromptId *string `field:"optional" json:"noteTakingAiPromptId" yaml:"noteTakingAiPromptId"`
}

