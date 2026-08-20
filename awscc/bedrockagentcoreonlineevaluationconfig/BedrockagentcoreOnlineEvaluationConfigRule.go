// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreonlineevaluationconfig


type BedrockagentcoreOnlineEvaluationConfigRule struct {
	// The configuration that controls what percentage of agent traces are sampled for evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_online_evaluation_config#sampling_config BedrockagentcoreOnlineEvaluationConfig#sampling_config}
	SamplingConfig *BedrockagentcoreOnlineEvaluationConfigRuleSamplingConfig `field:"required" json:"samplingConfig" yaml:"samplingConfig"`
	// The list of filters that determine which agent traces should be included in the evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_online_evaluation_config#filters BedrockagentcoreOnlineEvaluationConfig#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// The configuration that defines how agent sessions are detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_online_evaluation_config#session_config BedrockagentcoreOnlineEvaluationConfig#session_config}
	SessionConfig *BedrockagentcoreOnlineEvaluationConfigRuleSessionConfig `field:"optional" json:"sessionConfig" yaml:"sessionConfig"`
}

