// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable


type CleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholds struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_configured_table#allowed_aggregate_expression_type CleanroomsConfiguredTable#allowed_aggregate_expression_type}.
	AllowedAggregateExpressionType *string `field:"optional" json:"allowedAggregateExpressionType" yaml:"allowedAggregateExpressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_configured_table#identity_columns CleanroomsConfiguredTable#identity_columns}.
	IdentityColumns *[]*string `field:"optional" json:"identityColumns" yaml:"identityColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_configured_table#minimum_identity_count CleanroomsConfiguredTable#minimum_identity_count}.
	MinimumIdentityCount *float64 `field:"optional" json:"minimumIdentityCount" yaml:"minimumIdentityCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_configured_table#output_column_thresholds CleanroomsConfiguredTable#output_column_thresholds}.
	OutputColumnThresholds interface{} `field:"optional" json:"outputColumnThresholds" yaml:"outputColumnThresholds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_configured_table#type CleanroomsConfiguredTable#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

