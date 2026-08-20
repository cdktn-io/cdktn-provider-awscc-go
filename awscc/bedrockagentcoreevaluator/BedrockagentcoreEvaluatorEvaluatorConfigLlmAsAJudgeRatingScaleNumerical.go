// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleNumerical struct {
	// The description that explains what this numerical rating represents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_evaluator#definition BedrockagentcoreEvaluator#definition}
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// The label that describes this numerical rating option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_evaluator#label BedrockagentcoreEvaluator#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The numerical value for this rating scale option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_evaluator#value BedrockagentcoreEvaluator#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

