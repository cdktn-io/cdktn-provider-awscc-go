// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilescalculatedattributedefinition


type CustomerprofilesCalculatedAttributeDefinitionFilter struct {
	// Holds the list of Filter groups within the Filter definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/customerprofiles_calculated_attribute_definition#groups CustomerprofilesCalculatedAttributeDefinition#groups}
	Groups interface{} `field:"optional" json:"groups" yaml:"groups"`
	// Define whether to include or exclude objects for Calculated Attributed calculation that fit the filter groups criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/customerprofiles_calculated_attribute_definition#include CustomerprofilesCalculatedAttributeDefinition#include}
	Include *string `field:"optional" json:"include" yaml:"include"`
}

