// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeModelConfigResponsesEvaluatorModelConfig struct {
	// The maximum number of output tokens to generate, including visible output and reasoning tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_evaluator#max_output_tokens BedrockagentcoreEvaluator#max_output_tokens}
	MaxOutputTokens *float64 `field:"optional" json:"maxOutputTokens" yaml:"maxOutputTokens"`
	// The identifier of the model to use for evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_evaluator#model_id BedrockagentcoreEvaluator#model_id}
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// The reasoning configuration for reasoning models.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_evaluator#reasoning BedrockagentcoreEvaluator#reasoning}
	Reasoning *BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeModelConfigResponsesEvaluatorModelConfigReasoning `field:"optional" json:"reasoning" yaml:"reasoning"`
	// The sampling temperature between 0 and 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_evaluator#temperature BedrockagentcoreEvaluator#temperature}
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// The nucleus sampling probability mass between 0 and 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_evaluator#top_p BedrockagentcoreEvaluator#top_p}
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

