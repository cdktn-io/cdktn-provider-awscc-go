// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeModelConfig struct {
	// The configuration for using Amazon Bedrock models in evaluator assessments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_evaluator#bedrock_evaluator_model_config BedrockagentcoreEvaluator#bedrock_evaluator_model_config}
	BedrockEvaluatorModelConfig *BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeModelConfigBedrockEvaluatorModelConfig `field:"optional" json:"bedrockEvaluatorModelConfig" yaml:"bedrockEvaluatorModelConfig"`
	// The configuration for using OpenResponses-compatible models in evaluator assessments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_evaluator#responses_evaluator_model_config BedrockagentcoreEvaluator#responses_evaluator_model_config}
	ResponsesEvaluatorModelConfig *BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeModelConfigResponsesEvaluatorModelConfig `field:"optional" json:"responsesEvaluatorModelConfig" yaml:"responsesEvaluatorModelConfig"`
}

