// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationcentralizationrule


type ObservabilityadminOrganizationCentralizationRuleRuleDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/observabilityadmin_organization_centralization_rule#region ObservabilityadminOrganizationCentralizationRule#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/observabilityadmin_organization_centralization_rule#account ObservabilityadminOrganizationCentralizationRule#account}.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/observabilityadmin_organization_centralization_rule#destination_logs_configuration ObservabilityadminOrganizationCentralizationRule#destination_logs_configuration}.
	DestinationLogsConfiguration *ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfiguration `field:"optional" json:"destinationLogsConfiguration" yaml:"destinationLogsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/observabilityadmin_organization_centralization_rule#destination_metrics_configuration ObservabilityadminOrganizationCentralizationRule#destination_metrics_configuration}.
	DestinationMetricsConfiguration *ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationMetricsConfiguration `field:"optional" json:"destinationMetricsConfiguration" yaml:"destinationMetricsConfiguration"`
}

