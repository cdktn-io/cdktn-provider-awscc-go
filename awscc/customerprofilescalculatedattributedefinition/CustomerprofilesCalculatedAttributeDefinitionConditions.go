// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilescalculatedattributedefinition


type CustomerprofilesCalculatedAttributeDefinitionConditions struct {
	// The number of profile objects used for the calculated attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_calculated_attribute_definition#object_count CustomerprofilesCalculatedAttributeDefinition#object_count}
	ObjectCount *float64 `field:"optional" json:"objectCount" yaml:"objectCount"`
	// The relative time period over which data is included in the aggregation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_calculated_attribute_definition#range CustomerprofilesCalculatedAttributeDefinition#range}
	Range *CustomerprofilesCalculatedAttributeDefinitionConditionsRange `field:"optional" json:"range" yaml:"range"`
	// The threshold for the calculated attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_calculated_attribute_definition#threshold CustomerprofilesCalculatedAttributeDefinition#threshold}
	Threshold *CustomerprofilesCalculatedAttributeDefinitionConditionsThreshold `field:"optional" json:"threshold" yaml:"threshold"`
}

