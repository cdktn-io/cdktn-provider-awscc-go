// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable


type CleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputColumnThresholds struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanrooms_configured_table#minimum_identity_count CleanroomsConfiguredTable#minimum_identity_count}.
	MinimumIdentityCount *float64 `field:"optional" json:"minimumIdentityCount" yaml:"minimumIdentityCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cleanrooms_configured_table#output_column_name CleanroomsConfiguredTable#output_column_name}.
	OutputColumnName *string `field:"optional" json:"outputColumnName" yaml:"outputColumnName"`
}

