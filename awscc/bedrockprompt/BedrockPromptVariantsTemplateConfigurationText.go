// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsTemplateConfigurationText struct {
	// CachePointBlock.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_prompt#cache_point BedrockPrompt#cache_point}
	CachePoint *BedrockPromptVariantsTemplateConfigurationTextCachePoint `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// List of input variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_prompt#input_variables BedrockPrompt#input_variables}
	InputVariables interface{} `field:"optional" json:"inputVariables" yaml:"inputVariables"`
	// Prompt content for String prompt template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_prompt#text BedrockPrompt#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// The identifier for the S3 resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_prompt#text_s3_location BedrockPrompt#text_s3_location}
	TextS3Location *BedrockPromptVariantsTemplateConfigurationTextTextS3Location `field:"optional" json:"textS3Location" yaml:"textS3Location"`
}

