// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt


type BedrockPromptVariantsInferenceConfigurationText struct {
	// Maximum length of output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_prompt#max_tokens BedrockPrompt#max_tokens}
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// List of stop sequences.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_prompt#stop_sequences BedrockPrompt#stop_sequences}
	StopSequences *[]*string `field:"optional" json:"stopSequences" yaml:"stopSequences"`
	// Controls randomness, higher values increase diversity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_prompt#temperature BedrockPrompt#temperature}
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// Cumulative probability cutoff for token selection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_prompt#top_p BedrockPrompt#top_p}
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

