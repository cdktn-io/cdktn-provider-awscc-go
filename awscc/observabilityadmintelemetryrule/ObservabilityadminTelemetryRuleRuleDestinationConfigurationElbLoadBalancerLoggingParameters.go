// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule


type ObservabilityadminTelemetryRuleRuleDestinationConfigurationElbLoadBalancerLoggingParameters struct {
	// A delimiter to delineate log fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/observabilityadmin_telemetry_rule#field_delimiter ObservabilityadminTelemetryRule#field_delimiter}
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/observabilityadmin_telemetry_rule#output_format ObservabilityadminTelemetryRule#output_format}.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

