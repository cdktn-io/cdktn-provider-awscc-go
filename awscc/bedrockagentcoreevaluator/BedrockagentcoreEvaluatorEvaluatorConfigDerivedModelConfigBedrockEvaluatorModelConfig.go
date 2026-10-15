// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigBedrockEvaluatorModelConfig struct {
	// Additional model-specific request fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/bedrockagentcore_evaluator#additional_model_request_fields BedrockagentcoreEvaluator#additional_model_request_fields}
	AdditionalModelRequestFields *string `field:"optional" json:"additionalModelRequestFields" yaml:"additionalModelRequestFields"`
	// The inference configuration parameters that control model behavior during evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/bedrockagentcore_evaluator#inference_config BedrockagentcoreEvaluator#inference_config}
	InferenceConfig *BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigBedrockEvaluatorModelConfigInferenceConfig `field:"optional" json:"inferenceConfig" yaml:"inferenceConfig"`
	// The identifier of the Amazon Bedrock model to use for evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/bedrockagentcore_evaluator#model_id BedrockagentcoreEvaluator#model_id}
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
}

