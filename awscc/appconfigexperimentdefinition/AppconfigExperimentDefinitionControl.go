// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigexperimentdefinition


type AppconfigExperimentDefinitionControl struct {
	// Whether the flag is enabled for this variant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appconfig_experiment_definition#enabled AppconfigExperimentDefinition#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Traffic weight percentage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appconfig_experiment_definition#weight AppconfigExperimentDefinition#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Map of attribute name to attribute value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appconfig_experiment_definition#attribute_values AppconfigExperimentDefinition#attribute_values}
	AttributeValues interface{} `field:"optional" json:"attributeValues" yaml:"attributeValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appconfig_experiment_definition#description AppconfigExperimentDefinition#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

