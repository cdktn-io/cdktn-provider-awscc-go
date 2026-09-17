// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule


type ObservabilityadminTelemetryRuleRuleDestinationConfigurationCloudtrailParameters struct {
	// Create fine-grained selectors for AWS CloudTrail management and data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#advanced_event_selectors ObservabilityadminTelemetryRule#advanced_event_selectors}
	AdvancedEventSelectors interface{} `field:"optional" json:"advancedEventSelectors" yaml:"advancedEventSelectors"`
}

