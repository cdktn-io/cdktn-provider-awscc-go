// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationtelemetryrule


type ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationWafLoggingParametersLoggingFilter struct {
	// The behavior required of the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/observabilityadmin_organization_telemetry_rule#default_behavior ObservabilityadminOrganizationTelemetryRule#default_behavior}
	DefaultBehavior *string `field:"optional" json:"defaultBehavior" yaml:"defaultBehavior"`
	// A list of filters to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/observabilityadmin_organization_telemetry_rule#filters ObservabilityadminOrganizationTelemetryRule#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

