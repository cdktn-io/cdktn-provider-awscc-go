// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreonlineevaluationconfig


type BedrockagentcoreOnlineEvaluationConfigRuleFilters struct {
	// The key or field name to filter on within the agent trace data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_online_evaluation_config#key BedrockagentcoreOnlineEvaluationConfig#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The comparison operator to use for filtering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_online_evaluation_config#operator BedrockagentcoreOnlineEvaluationConfig#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The value used in filter comparisons.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_online_evaluation_config#value BedrockagentcoreOnlineEvaluationConfig#value}
	Value *BedrockagentcoreOnlineEvaluationConfigRuleFiltersValue `field:"optional" json:"value" yaml:"value"`
}

