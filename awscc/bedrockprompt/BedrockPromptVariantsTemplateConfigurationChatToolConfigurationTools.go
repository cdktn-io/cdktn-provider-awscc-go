// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsTemplateConfigurationChatToolConfigurationTools struct {
	// CachePointBlock.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_prompt#cache_point BedrockPrompt#cache_point}
	CachePoint *BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsCachePoint `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// Tool specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_prompt#tool_spec BedrockPrompt#tool_spec}
	ToolSpec *BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsToolSpec `field:"optional" json:"toolSpec" yaml:"toolSpec"`
}

