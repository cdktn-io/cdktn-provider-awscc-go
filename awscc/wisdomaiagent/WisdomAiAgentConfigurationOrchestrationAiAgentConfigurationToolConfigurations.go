// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#annotations WisdomAiAgent#annotations}.
	Annotations *string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#description WisdomAiAgent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#input_schema WisdomAiAgent#input_schema}.
	InputSchema *string `field:"optional" json:"inputSchema" yaml:"inputSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#instruction WisdomAiAgent#instruction}.
	Instruction *WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsInstruction `field:"optional" json:"instruction" yaml:"instruction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#output_filters WisdomAiAgent#output_filters}.
	OutputFilters interface{} `field:"optional" json:"outputFilters" yaml:"outputFilters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#output_schema WisdomAiAgent#output_schema}.
	OutputSchema *string `field:"optional" json:"outputSchema" yaml:"outputSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#override_input_values WisdomAiAgent#override_input_values}.
	OverrideInputValues interface{} `field:"optional" json:"overrideInputValues" yaml:"overrideInputValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#title WisdomAiAgent#title}.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#tool_id WisdomAiAgent#tool_id}.
	ToolId *string `field:"optional" json:"toolId" yaml:"toolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#tool_name WisdomAiAgent#tool_name}.
	ToolName *string `field:"optional" json:"toolName" yaml:"toolName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#tool_type WisdomAiAgent#tool_type}.
	ToolType *string `field:"optional" json:"toolType" yaml:"toolType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#user_interaction_configuration WisdomAiAgent#user_interaction_configuration}.
	UserInteractionConfiguration *WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsUserInteractionConfiguration `field:"optional" json:"userInteractionConfiguration" yaml:"userInteractionConfiguration"`
}

