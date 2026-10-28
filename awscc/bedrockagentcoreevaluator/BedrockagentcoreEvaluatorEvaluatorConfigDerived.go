// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfigDerived struct {
	// The identifier of the base evaluator whose logic to run (a Builtin.* or ThirdParty.* evaluator).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrockagentcore_evaluator#base_evaluator_id BedrockagentcoreEvaluator#base_evaluator_id}
	BaseEvaluatorId *string `field:"optional" json:"baseEvaluatorId" yaml:"baseEvaluatorId"`
	// The model configuration that specifies which foundation model to use for evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrockagentcore_evaluator#model_config BedrockagentcoreEvaluator#model_config}
	ModelConfig *BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfig `field:"optional" json:"modelConfig" yaml:"modelConfig"`
}

