// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreonlineevaluationconfig


type BedrockagentcoreOnlineEvaluationConfigRuleSamplingConfig struct {
	// The percentage of agent traces to sample for evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_online_evaluation_config#sampling_percentage BedrockagentcoreOnlineEvaluationConfig#sampling_percentage}
	SamplingPercentage *float64 `field:"required" json:"samplingPercentage" yaml:"samplingPercentage"`
}

