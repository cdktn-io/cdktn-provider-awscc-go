// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreonlineevaluationconfig


type BedrockagentcoreOnlineEvaluationConfigEvaluators struct {
	// The ID of the evaluator: a built-in evaluator ID (Builtin.*), a managed third-party evaluator ID (ThirdParty.<provider>.<name>), or a custom evaluator ID. For a custom evaluator in the same template, use Fn::GetAtt with EvaluatorId. Ref returns the evaluator ARN, which this property does not accept.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/bedrockagentcore_online_evaluation_config#evaluator_id BedrockagentcoreOnlineEvaluationConfig#evaluator_id}
	EvaluatorId *string `field:"optional" json:"evaluatorId" yaml:"evaluatorId"`
}

