// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule


type ObservabilityadminTelemetryRuleRuleDestinationConfigurationWafLoggingParametersLoggingFilter struct {
	// The behavior required of the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_telemetry_rule#default_behavior ObservabilityadminTelemetryRule#default_behavior}
	DefaultBehavior *string `field:"optional" json:"defaultBehavior" yaml:"defaultBehavior"`
	// A list of filters to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_telemetry_rule#filters ObservabilityadminTelemetryRule#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

