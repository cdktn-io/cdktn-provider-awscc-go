// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationcentralizationrule


type ObservabilityadminOrganizationCentralizationRuleRuleSourceSourceLogsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_organization_centralization_rule#data_source_selection_criteria ObservabilityadminOrganizationCentralizationRule#data_source_selection_criteria}.
	DataSourceSelectionCriteria *string `field:"optional" json:"dataSourceSelectionCriteria" yaml:"dataSourceSelectionCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_organization_centralization_rule#encrypted_log_group_strategy ObservabilityadminOrganizationCentralizationRule#encrypted_log_group_strategy}.
	EncryptedLogGroupStrategy *string `field:"optional" json:"encryptedLogGroupStrategy" yaml:"encryptedLogGroupStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_organization_centralization_rule#log_group_selection_criteria ObservabilityadminOrganizationCentralizationRule#log_group_selection_criteria}.
	LogGroupSelectionCriteria *string `field:"optional" json:"logGroupSelectionCriteria" yaml:"logGroupSelectionCriteria"`
}

