// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsintermediatetable


type CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomComparisonControls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_intermediate_table#allowed_column_comparison_columns CleanroomsIntermediateTable#allowed_column_comparison_columns}.
	AllowedColumnComparisonColumns *[]*string `field:"optional" json:"allowedColumnComparisonColumns" yaml:"allowedColumnComparisonColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cleanrooms_intermediate_table#allowed_literal_comparison_columns CleanroomsIntermediateTable#allowed_literal_comparison_columns}.
	AllowedLiteralComparisonColumns *[]*string `field:"optional" json:"allowedLiteralComparisonColumns" yaml:"allowedLiteralComparisonColumns"`
}

