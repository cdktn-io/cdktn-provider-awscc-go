// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilescalculatedattributedefinition


type CustomerprofilesCalculatedAttributeDefinitionFilterGroups struct {
	// Object that holds the attributes to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/customerprofiles_calculated_attribute_definition#dimensions CustomerprofilesCalculatedAttributeDefinition#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The type of logical relationship between the dimensions of the Filter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/customerprofiles_calculated_attribute_definition#type CustomerprofilesCalculatedAttributeDefinition#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

