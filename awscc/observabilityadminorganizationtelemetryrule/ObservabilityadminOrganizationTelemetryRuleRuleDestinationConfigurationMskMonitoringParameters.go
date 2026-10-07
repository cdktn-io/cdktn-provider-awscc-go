// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationtelemetryrule


type ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationMskMonitoringParameters struct {
	// The level of enhanced monitoring for the MSK cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/observabilityadmin_organization_telemetry_rule#enhanced_monitoring ObservabilityadminOrganizationTelemetryRule#enhanced_monitoring}
	EnhancedMonitoring *string `field:"optional" json:"enhancedMonitoring" yaml:"enhancedMonitoring"`
}

