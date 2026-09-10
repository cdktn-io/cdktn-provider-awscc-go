// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationtelemetryrule


type ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationLogDeliveryParameters struct {
	// Types of logs to deliver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/observabilityadmin_organization_telemetry_rule#log_types ObservabilityadminOrganizationTelemetryRule#log_types}
	LogTypes *[]*string `field:"optional" json:"logTypes" yaml:"logTypes"`
}

