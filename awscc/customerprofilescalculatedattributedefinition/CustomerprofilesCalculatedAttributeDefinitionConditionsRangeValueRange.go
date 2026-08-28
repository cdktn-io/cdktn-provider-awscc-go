// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilescalculatedattributedefinition


type CustomerprofilesCalculatedAttributeDefinitionConditionsRangeValueRange struct {
	// The ending point for this range.
	//
	// Positive numbers indicate how many days in the past data should be included, and negative numbers indicate how many days in the future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_calculated_attribute_definition#end CustomerprofilesCalculatedAttributeDefinition#end}
	End *float64 `field:"optional" json:"end" yaml:"end"`
	// The starting point for this range.
	//
	// Positive numbers indicate how many days in the past data should be included, and negative numbers indicate how many days in the future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_calculated_attribute_definition#start CustomerprofilesCalculatedAttributeDefinition#start}
	Start *float64 `field:"optional" json:"start" yaml:"start"`
}

