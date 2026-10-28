// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrockagentcore_evaluator#categorical BedrockagentcoreEvaluator#categorical}.
	Categorical interface{} `field:"optional" json:"categorical" yaml:"categorical"`
	// The dimensions an LLM-as-a-Judge evaluator scores in a single judge call.
	//
	// Each dimension has its own instruction and rating scale.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrockagentcore_evaluator#multi_dimensional BedrockagentcoreEvaluator#multi_dimensional}
	MultiDimensional interface{} `field:"optional" json:"multiDimensional" yaml:"multiDimensional"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrockagentcore_evaluator#numerical BedrockagentcoreEvaluator#numerical}.
	Numerical interface{} `field:"optional" json:"numerical" yaml:"numerical"`
}

