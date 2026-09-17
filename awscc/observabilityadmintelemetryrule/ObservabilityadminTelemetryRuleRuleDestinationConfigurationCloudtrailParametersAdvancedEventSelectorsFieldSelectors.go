// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule


type ObservabilityadminTelemetryRuleRuleDestinationConfigurationCloudtrailParametersAdvancedEventSelectorsFieldSelectors struct {
	// An operator that includes events that match the last few characters of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#ends_with ObservabilityadminTelemetryRule#ends_with}
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// An operator that includes events that match the exact value of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#equals ObservabilityadminTelemetryRule#equals}
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// A field in a CloudTrail event record on which to filter events to be logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#field ObservabilityadminTelemetryRule#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// An operator that excludes events that match the last few characters of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#not_ends_with ObservabilityadminTelemetryRule#not_ends_with}
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// An operator that excludes events that match the exact value of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#not_equals ObservabilityadminTelemetryRule#not_equals}
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// An operator that excludes events that match the first few characters of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#not_starts_with ObservabilityadminTelemetryRule#not_starts_with}
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// An operator that includes events that match the first few characters of the event record field specified as the value of Field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/observabilityadmin_telemetry_rule#starts_with ObservabilityadminTelemetryRule#starts_with}
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

