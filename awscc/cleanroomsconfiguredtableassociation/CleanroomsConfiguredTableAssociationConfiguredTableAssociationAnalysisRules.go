// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtableassociation


type CleanroomsConfiguredTableAssociationConfiguredTableAssociationAnalysisRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cleanrooms_configured_table_association#policy CleanroomsConfiguredTableAssociation#policy}.
	Policy *CleanroomsConfiguredTableAssociationConfiguredTableAssociationAnalysisRulesPolicy `field:"optional" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cleanrooms_configured_table_association#type CleanroomsConfiguredTableAssociation#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

