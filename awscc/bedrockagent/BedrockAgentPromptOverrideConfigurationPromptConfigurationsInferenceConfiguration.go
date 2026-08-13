// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent


type BedrockAgentPromptOverrideConfigurationPromptConfigurationsInferenceConfiguration struct {
	// Maximum length of output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_agent#maximum_length BedrockAgent#maximum_length}
	MaximumLength *float64 `field:"optional" json:"maximumLength" yaml:"maximumLength"`
	// List of stop sequences.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_agent#stop_sequences BedrockAgent#stop_sequences}
	StopSequences *[]*string `field:"optional" json:"stopSequences" yaml:"stopSequences"`
	// Controls randomness, higher values increase diversity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_agent#temperature BedrockAgent#temperature}
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// Sample from the k most likely next tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_agent#top_k BedrockAgent#top_k}
	TopK *float64 `field:"optional" json:"topK" yaml:"topK"`
	// Cumulative probability cutoff for token selection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_agent#top_p BedrockAgent#top_p}
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

