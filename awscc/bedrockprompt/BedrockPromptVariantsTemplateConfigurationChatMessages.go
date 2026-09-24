// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsTemplateConfigurationChatMessages struct {
	// List of Content Blocks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#content BedrockPrompt#content}
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// Conversation roles for the chat prompt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_prompt#role BedrockPrompt#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

