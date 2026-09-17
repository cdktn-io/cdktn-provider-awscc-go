// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule


type ObservabilityadminTelemetryRuleRule struct {
	// Resource Type associated with the Telemetry Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#resource_type ObservabilityadminTelemetryRule#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Telemetry Type associated with the Telemetry Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#telemetry_type ObservabilityadminTelemetryRule#telemetry_type}
	TelemetryType *string `field:"required" json:"telemetryType" yaml:"telemetryType"`
	// When true, configuration drift in managed telemetry resources will be detected and remediated for resource-level fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#allow_field_updates ObservabilityadminTelemetryRule#allow_field_updates}
	AllowFieldUpdates interface{} `field:"optional" json:"allowFieldUpdates" yaml:"allowFieldUpdates"`
	// When true, the rule is replicated to all supported regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#all_regions ObservabilityadminTelemetryRule#all_regions}
	AllRegions interface{} `field:"optional" json:"allRegions" yaml:"allRegions"`
	// The destination configuration for telemetry data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#destination_configuration ObservabilityadminTelemetryRule#destination_configuration}
	DestinationConfiguration *ObservabilityadminTelemetryRuleRuleDestinationConfiguration `field:"optional" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// List of AWS region codes where the rule should be replicated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#regions ObservabilityadminTelemetryRule#regions}
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Selection Criteria on resource level for rule application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#selection_criteria ObservabilityadminTelemetryRule#selection_criteria}
	SelectionCriteria *string `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
	// The telemetry source types for a telemetry rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#telemetry_source_types ObservabilityadminTelemetryRule#telemetry_source_types}
	TelemetrySourceTypes *[]*string `field:"optional" json:"telemetrySourceTypes" yaml:"telemetrySourceTypes"`
}

