// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsintermediatetable


type CleanroomsIntermediateTableAnalysisRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanrooms_intermediate_table#policy CleanroomsIntermediateTable#policy}.
	Policy *CleanroomsIntermediateTableAnalysisRulesPolicy `field:"optional" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cleanrooms_intermediate_table#type CleanroomsIntermediateTable#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

