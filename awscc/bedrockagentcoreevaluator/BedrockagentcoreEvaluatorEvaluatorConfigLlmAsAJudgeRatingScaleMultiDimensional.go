// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensional struct {
	// The instruction to the evaluator model describing what to judge for this dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrockagentcore_evaluator#dimension_instruction BedrockagentcoreEvaluator#dimension_instruction}
	DimensionInstruction *string `field:"optional" json:"dimensionInstruction" yaml:"dimensionInstruction"`
	// The name of the dimension. Must be unique within the evaluator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrockagentcore_evaluator#dimension_name BedrockagentcoreEvaluator#dimension_name}
	DimensionName *string `field:"optional" json:"dimensionName" yaml:"dimensionName"`
	// The rating scale for one dimension of an LLM-as-a-Judge evaluator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrockagentcore_evaluator#dimension_rating_scale BedrockagentcoreEvaluator#dimension_rating_scale}
	DimensionRatingScale *BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalDimensionRatingScale `field:"optional" json:"dimensionRatingScale" yaml:"dimensionRatingScale"`
}

