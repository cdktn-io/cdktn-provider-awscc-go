// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreonlineevaluationconfig


type BedrockagentcoreOnlineEvaluationConfigRuleFiltersValue struct {
	// The boolean value for true/false filtering conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_online_evaluation_config#boolean_value BedrockagentcoreOnlineEvaluationConfig#boolean_value}
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// The numeric value for numerical filtering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_online_evaluation_config#double_value BedrockagentcoreOnlineEvaluationConfig#double_value}
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// The string value for text-based filtering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_online_evaluation_config#string_value BedrockagentcoreOnlineEvaluationConfig#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

