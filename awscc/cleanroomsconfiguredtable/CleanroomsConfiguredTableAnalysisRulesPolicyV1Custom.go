// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable


type CleanroomsConfiguredTableAnalysisRulesPolicyV1Custom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table#additional_analyses CleanroomsConfiguredTable#additional_analyses}.
	AdditionalAnalyses *string `field:"optional" json:"additionalAnalyses" yaml:"additionalAnalyses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table#allowed_analyses CleanroomsConfiguredTable#allowed_analyses}.
	AllowedAnalyses *[]*string `field:"optional" json:"allowedAnalyses" yaml:"allowedAnalyses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table#allowed_analysis_providers CleanroomsConfiguredTable#allowed_analysis_providers}.
	AllowedAnalysisProviders *[]*string `field:"optional" json:"allowedAnalysisProviders" yaml:"allowedAnalysisProviders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table#differential_privacy CleanroomsConfiguredTable#differential_privacy}.
	DifferentialPrivacy *CleanroomsConfiguredTableAnalysisRulesPolicyV1CustomDifferentialPrivacy `field:"optional" json:"differentialPrivacy" yaml:"differentialPrivacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table#disallowed_output_columns CleanroomsConfiguredTable#disallowed_output_columns}.
	DisallowedOutputColumns *[]*string `field:"optional" json:"disallowedOutputColumns" yaml:"disallowedOutputColumns"`
}

