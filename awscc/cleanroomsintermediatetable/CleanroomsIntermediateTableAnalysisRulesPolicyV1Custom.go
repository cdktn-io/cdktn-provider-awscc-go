// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsintermediatetable


type CleanroomsIntermediateTableAnalysisRulesPolicyV1Custom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_intermediate_table#additional_analyses CleanroomsIntermediateTable#additional_analyses}.
	AdditionalAnalyses *string `field:"optional" json:"additionalAnalyses" yaml:"additionalAnalyses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_intermediate_table#allowed_analyses CleanroomsIntermediateTable#allowed_analyses}.
	AllowedAnalyses *[]*string `field:"optional" json:"allowedAnalyses" yaml:"allowedAnalyses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_intermediate_table#allowed_analysis_providers CleanroomsIntermediateTable#allowed_analysis_providers}.
	AllowedAnalysisProviders *[]*string `field:"optional" json:"allowedAnalysisProviders" yaml:"allowedAnalysisProviders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_intermediate_table#allowed_result_receivers CleanroomsIntermediateTable#allowed_result_receivers}.
	AllowedResultReceivers *[]*string `field:"optional" json:"allowedResultReceivers" yaml:"allowedResultReceivers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_intermediate_table#differential_privacy CleanroomsIntermediateTable#differential_privacy}.
	DifferentialPrivacy *CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomDifferentialPrivacy `field:"optional" json:"differentialPrivacy" yaml:"differentialPrivacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_intermediate_table#disallowed_output_columns CleanroomsIntermediateTable#disallowed_output_columns}.
	DisallowedOutputColumns *[]*string `field:"optional" json:"disallowedOutputColumns" yaml:"disallowedOutputColumns"`
}

