// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsTemplateConfiguration struct {
	// Configuration for chat prompt template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_prompt#chat BedrockPrompt#chat}
	Chat *BedrockPromptVariantsTemplateConfigurationChat `field:"optional" json:"chat" yaml:"chat"`
	// Configuration for text prompt template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_prompt#text BedrockPrompt#text}
	Text *BedrockPromptVariantsTemplateConfigurationText `field:"optional" json:"text" yaml:"text"`
}

