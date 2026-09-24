// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsTemplateConfigurationChat struct {
	// List of input variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#input_variables BedrockPrompt#input_variables}
	InputVariables interface{} `field:"optional" json:"inputVariables" yaml:"inputVariables"`
	// List of messages for chat prompt template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#messages BedrockPrompt#messages}
	Messages interface{} `field:"optional" json:"messages" yaml:"messages"`
	// Configuration for chat prompt template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#system BedrockPrompt#system}
	SystemAttribute interface{} `field:"optional" json:"systemAttribute" yaml:"systemAttribute"`
	// Tool configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#tool_configuration BedrockPrompt#tool_configuration}
	ToolConfiguration *BedrockPromptVariantsTemplateConfigurationChatToolConfiguration `field:"optional" json:"toolConfiguration" yaml:"toolConfiguration"`
}

