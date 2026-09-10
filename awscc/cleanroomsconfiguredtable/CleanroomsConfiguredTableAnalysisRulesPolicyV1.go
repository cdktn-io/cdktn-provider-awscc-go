// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable


type CleanroomsConfiguredTableAnalysisRulesPolicyV1 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cleanrooms_configured_table#aggregation CleanroomsConfiguredTable#aggregation}.
	Aggregation *CleanroomsConfiguredTableAnalysisRulesPolicyV1Aggregation `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cleanrooms_configured_table#custom CleanroomsConfiguredTable#custom}.
	Custom *CleanroomsConfiguredTableAnalysisRulesPolicyV1Custom `field:"optional" json:"custom" yaml:"custom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cleanrooms_configured_table#list CleanroomsConfiguredTable#list}.
	List *CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStruct `field:"optional" json:"list" yaml:"list"`
}

