// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable


type CleanroomsConfiguredTableAnalysisRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanrooms_configured_table#policy CleanroomsConfiguredTable#policy}.
	Policy *CleanroomsConfiguredTableAnalysisRulesPolicy `field:"optional" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanrooms_configured_table#type CleanroomsConfiguredTable#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

