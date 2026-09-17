// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeModelConfigBedrockEvaluatorModelConfigInferenceConfig struct {
	// The maximum number of tokens to generate in the model response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_evaluator#max_tokens BedrockagentcoreEvaluator#max_tokens}
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// The temperature value that controls randomness in the model's responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_evaluator#temperature BedrockagentcoreEvaluator#temperature}
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// The top-p sampling parameter that controls the diversity of the model's responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_evaluator#top_p BedrockagentcoreEvaluator#top_p}
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

