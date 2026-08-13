// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfig struct {
	// The configuration for code-based evaluation using a Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_evaluator#code_based BedrockagentcoreEvaluator#code_based}
	CodeBased *BedrockagentcoreEvaluatorEvaluatorConfigCodeBased `field:"optional" json:"codeBased" yaml:"codeBased"`
	// The configuration for LLM-as-a-Judge evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_evaluator#llm_as_a_judge BedrockagentcoreEvaluator#llm_as_a_judge}
	LlmAsAJudge *BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudge `field:"optional" json:"llmAsAJudge" yaml:"llmAsAJudge"`
}

