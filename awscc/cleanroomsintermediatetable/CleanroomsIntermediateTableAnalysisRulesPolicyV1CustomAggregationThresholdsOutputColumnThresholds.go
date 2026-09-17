// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsintermediatetable


type CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputColumnThresholds struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanrooms_intermediate_table#minimum_identity_count CleanroomsIntermediateTable#minimum_identity_count}.
	MinimumIdentityCount *float64 `field:"optional" json:"minimumIdentityCount" yaml:"minimumIdentityCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanrooms_intermediate_table#output_column_name CleanroomsIntermediateTable#output_column_name}.
	OutputColumnName *string `field:"optional" json:"outputColumnName" yaml:"outputColumnName"`
}

