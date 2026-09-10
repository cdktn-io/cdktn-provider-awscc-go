// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationcentralizationrule


type ObservabilityadminOrganizationCentralizationRuleRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/observabilityadmin_organization_centralization_rule#destination ObservabilityadminOrganizationCentralizationRule#destination}.
	Destination *ObservabilityadminOrganizationCentralizationRuleRuleDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/observabilityadmin_organization_centralization_rule#source ObservabilityadminOrganizationCentralizationRule#source}.
	Source *ObservabilityadminOrganizationCentralizationRuleRuleSource `field:"required" json:"source" yaml:"source"`
}

