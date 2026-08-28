// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationtelemetryrule


type ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationWafLoggingParametersLoggingFilterFilters struct {
	// The behavior required of the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#behavior ObservabilityadminOrganizationTelemetryRule#behavior}
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// A list of conditions for a filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#conditions ObservabilityadminOrganizationTelemetryRule#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The requirement portion of the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#requirement ObservabilityadminOrganizationTelemetryRule#requirement}
	Requirement *string `field:"optional" json:"requirement" yaml:"requirement"`
}

