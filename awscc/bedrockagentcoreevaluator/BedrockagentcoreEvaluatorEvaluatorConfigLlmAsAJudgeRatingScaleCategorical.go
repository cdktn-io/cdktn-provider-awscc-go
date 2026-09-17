// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleCategorical struct {
	// The description that explains what this categorical rating represents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_evaluator#definition BedrockagentcoreEvaluator#definition}
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// The label of this categorical rating option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_evaluator#label BedrockagentcoreEvaluator#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}

