// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalDimensionRatingScale struct {
	// The categorical rating scale for this dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrockagentcore_evaluator#categorical BedrockagentcoreEvaluator#categorical}
	Categorical interface{} `field:"optional" json:"categorical" yaml:"categorical"`
	// The numerical rating scale for this dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrockagentcore_evaluator#numerical BedrockagentcoreEvaluator#numerical}
	Numerical interface{} `field:"optional" json:"numerical" yaml:"numerical"`
}

