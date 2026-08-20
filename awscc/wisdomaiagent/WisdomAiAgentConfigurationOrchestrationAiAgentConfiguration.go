// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationOrchestrationAiAgentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#connect_instance_arn WisdomAiAgent#connect_instance_arn}.
	ConnectInstanceArn *string `field:"optional" json:"connectInstanceArn" yaml:"connectInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#locale WisdomAiAgent#locale}.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#orchestration_ai_guardrail_id WisdomAiAgent#orchestration_ai_guardrail_id}.
	OrchestrationAiGuardrailId *string `field:"optional" json:"orchestrationAiGuardrailId" yaml:"orchestrationAiGuardrailId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#orchestration_ai_prompt_id WisdomAiAgent#orchestration_ai_prompt_id}.
	OrchestrationAiPromptId *string `field:"optional" json:"orchestrationAiPromptId" yaml:"orchestrationAiPromptId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#tool_configurations WisdomAiAgent#tool_configurations}.
	ToolConfigurations interface{} `field:"optional" json:"toolConfigurations" yaml:"toolConfigurations"`
}

