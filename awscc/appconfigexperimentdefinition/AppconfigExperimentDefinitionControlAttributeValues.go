// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigexperimentdefinition


type AppconfigExperimentDefinitionControlAttributeValues struct {
	// A boolean value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appconfig_experiment_definition#boolean_value AppconfigExperimentDefinition#boolean_value}
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// An array of numeric values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appconfig_experiment_definition#number_array AppconfigExperimentDefinition#number_array}
	NumberArray *[]*float64 `field:"optional" json:"numberArray" yaml:"numberArray"`
	// A numeric value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appconfig_experiment_definition#number_value AppconfigExperimentDefinition#number_value}
	NumberValue *float64 `field:"optional" json:"numberValue" yaml:"numberValue"`
	// An array of string values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appconfig_experiment_definition#string_array AppconfigExperimentDefinition#string_array}
	StringArray *[]*string `field:"optional" json:"stringArray" yaml:"stringArray"`
	// A string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appconfig_experiment_definition#string_value AppconfigExperimentDefinition#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

