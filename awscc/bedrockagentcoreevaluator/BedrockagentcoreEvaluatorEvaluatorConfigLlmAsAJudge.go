// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudge struct {
	// The evaluation instructions that guide the language model in assessing agent performance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_evaluator#instructions BedrockagentcoreEvaluator#instructions}
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// The model configuration that specifies which foundation model to use for evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_evaluator#model_config BedrockagentcoreEvaluator#model_config}
	ModelConfig *BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeModelConfig `field:"optional" json:"modelConfig" yaml:"modelConfig"`
	// The rating scale that defines how evaluators should score agent performance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_evaluator#rating_scale BedrockagentcoreEvaluator#rating_scale}
	RatingScale *BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScale `field:"optional" json:"ratingScale" yaml:"ratingScale"`
}

