// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilescalculatedattributedefinition


type CustomerprofilesCalculatedAttributeDefinitionAttributeDetails struct {
	// A list of attribute items specified in the mathematical expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_calculated_attribute_definition#attributes CustomerprofilesCalculatedAttributeDefinition#attributes}
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// Mathematical expression that is performed on attribute items provided in the attribute list.
	//
	// Each element in the expression should follow the structure of "{ObjectTypeName.AttributeName}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_calculated_attribute_definition#expression CustomerprofilesCalculatedAttributeDefinition#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

