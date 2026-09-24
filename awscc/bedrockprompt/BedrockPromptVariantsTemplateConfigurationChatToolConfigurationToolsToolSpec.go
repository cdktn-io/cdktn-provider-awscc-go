// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsToolSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#description BedrockPrompt#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tool input schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#input_schema BedrockPrompt#input_schema}
	InputSchema *BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsToolSpecInputSchema `field:"optional" json:"inputSchema" yaml:"inputSchema"`
	// Tool name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#name BedrockPrompt#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

