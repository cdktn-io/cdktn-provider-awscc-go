// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable


type CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputConstraints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_configured_table#column_name CleanroomsConfiguredTable#column_name}.
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_configured_table#minimum CleanroomsConfiguredTable#minimum}.
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_configured_table#type CleanroomsConfiguredTable#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

