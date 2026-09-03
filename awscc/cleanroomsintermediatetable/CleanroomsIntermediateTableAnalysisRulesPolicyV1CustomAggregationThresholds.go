// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsintermediatetable


type CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholds struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_intermediate_table#allowed_aggregate_expression_type CleanroomsIntermediateTable#allowed_aggregate_expression_type}.
	AllowedAggregateExpressionType *string `field:"optional" json:"allowedAggregateExpressionType" yaml:"allowedAggregateExpressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_intermediate_table#identity_columns CleanroomsIntermediateTable#identity_columns}.
	IdentityColumns *[]*string `field:"optional" json:"identityColumns" yaml:"identityColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_intermediate_table#minimum_identity_count CleanroomsIntermediateTable#minimum_identity_count}.
	MinimumIdentityCount *float64 `field:"optional" json:"minimumIdentityCount" yaml:"minimumIdentityCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_intermediate_table#output_column_thresholds CleanroomsIntermediateTable#output_column_thresholds}.
	OutputColumnThresholds interface{} `field:"optional" json:"outputColumnThresholds" yaml:"outputColumnThresholds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_intermediate_table#type CleanroomsIntermediateTable#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

