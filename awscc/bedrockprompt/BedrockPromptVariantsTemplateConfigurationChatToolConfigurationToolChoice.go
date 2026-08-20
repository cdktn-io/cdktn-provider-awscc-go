// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoice struct {
	// Any Tool choice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_prompt#any BedrockPrompt#any}
	Any *string `field:"optional" json:"any" yaml:"any"`
	// Auto Tool choice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_prompt#auto BedrockPrompt#auto}
	Auto *string `field:"optional" json:"auto" yaml:"auto"`
	// Specific Tool choice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_prompt#tool BedrockPrompt#tool}
	Tool *BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceTool `field:"optional" json:"tool" yaml:"tool"`
}

