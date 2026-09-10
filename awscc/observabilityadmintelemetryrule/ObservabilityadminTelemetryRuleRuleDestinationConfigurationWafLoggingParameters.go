// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule


type ObservabilityadminTelemetryRuleRuleDestinationConfigurationWafLoggingParameters struct {
	// Default handling for logs that don't match any of the specified filtering conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/observabilityadmin_telemetry_rule#logging_filter ObservabilityadminTelemetryRule#logging_filter}
	LoggingFilter *ObservabilityadminTelemetryRuleRuleDestinationConfigurationWafLoggingParametersLoggingFilter `field:"optional" json:"loggingFilter" yaml:"loggingFilter"`
	// The type of logs to generate for WAF.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/observabilityadmin_telemetry_rule#log_type ObservabilityadminTelemetryRule#log_type}
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
	// Fields not to be included in the logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/observabilityadmin_telemetry_rule#redacted_fields ObservabilityadminTelemetryRule#redacted_fields}
	RedactedFields interface{} `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

