// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtableassociation


type CleanroomsConfiguredTableAssociationConfiguredTableAssociationAnalysisRulesPolicyV1Aggregation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_configured_table_association#allowed_additional_analyses CleanroomsConfiguredTableAssociation#allowed_additional_analyses}.
	AllowedAdditionalAnalyses *[]*string `field:"optional" json:"allowedAdditionalAnalyses" yaml:"allowedAdditionalAnalyses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_configured_table_association#allowed_result_receivers CleanroomsConfiguredTableAssociation#allowed_result_receivers}.
	AllowedResultReceivers *[]*string `field:"optional" json:"allowedResultReceivers" yaml:"allowedResultReceivers"`
}

