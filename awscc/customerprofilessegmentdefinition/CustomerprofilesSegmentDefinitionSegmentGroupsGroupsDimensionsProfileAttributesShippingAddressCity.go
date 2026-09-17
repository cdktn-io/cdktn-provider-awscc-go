// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition


type CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCity struct {
	// The type of segment dimension to use for a string dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/customerprofiles_segment_definition#dimension_type CustomerprofilesSegmentDefinition#dimension_type}
	DimensionType *string `field:"optional" json:"dimensionType" yaml:"dimensionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/customerprofiles_segment_definition#values CustomerprofilesSegmentDefinition#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

