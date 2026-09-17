// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsTemplateConfigurationChatToolConfiguration struct {
	// Tool choice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_prompt#tool_choice BedrockPrompt#tool_choice}
	ToolChoice *BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoice `field:"optional" json:"toolChoice" yaml:"toolChoice"`
	// List of Tools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_prompt#tools BedrockPrompt#tools}
	Tools interface{} `field:"optional" json:"tools" yaml:"tools"`
}

