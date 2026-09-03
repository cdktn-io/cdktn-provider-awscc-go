// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition


type CustomerprofilesSegmentDefinitionSegmentSortAttributes struct {
	// The data type of the sort attribute (e.g., string, number, date).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#data_type CustomerprofilesSegmentDefinition#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// The name of the attribute to sort by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#name CustomerprofilesSegmentDefinition#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The sort order for the attribute (ascending or descending).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#order CustomerprofilesSegmentDefinition#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// The type of attribute (e.g., profile, calculated).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#type CustomerprofilesSegmentDefinition#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

