// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsGenAiResource struct {
	// Target Agent to invoke with Prompt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_prompt#agent BedrockPrompt#agent}
	Agent *BedrockPromptVariantsGenAiResourceAgent `field:"optional" json:"agent" yaml:"agent"`
}

