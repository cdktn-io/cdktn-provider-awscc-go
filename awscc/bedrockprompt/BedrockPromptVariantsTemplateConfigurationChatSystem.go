// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsTemplateConfigurationChatSystem struct {
	// CachePointBlock.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_prompt#cache_point BedrockPrompt#cache_point}
	CachePoint *BedrockPromptVariantsTemplateConfigurationChatSystemCachePoint `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// Configuration for chat prompt template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_prompt#text BedrockPrompt#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}

