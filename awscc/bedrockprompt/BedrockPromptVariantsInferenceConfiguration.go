// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsInferenceConfiguration struct {
	// Prompt model inference configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_prompt#text BedrockPrompt#text}
	Text *BedrockPromptVariantsInferenceConfigurationText `field:"optional" json:"text" yaml:"text"`
}

