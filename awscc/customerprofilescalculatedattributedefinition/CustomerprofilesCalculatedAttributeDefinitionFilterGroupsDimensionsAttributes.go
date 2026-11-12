// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilescalculatedattributedefinition


type CustomerprofilesCalculatedAttributeDefinitionFilterGroupsDimensionsAttributes struct {
	// The action to filter with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/customerprofiles_calculated_attribute_definition#dimension_type CustomerprofilesCalculatedAttributeDefinition#dimension_type}
	DimensionType *string `field:"optional" json:"dimensionType" yaml:"dimensionType"`
	// The values to apply the DimensionType on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/customerprofiles_calculated_attribute_definition#values CustomerprofilesCalculatedAttributeDefinition#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

