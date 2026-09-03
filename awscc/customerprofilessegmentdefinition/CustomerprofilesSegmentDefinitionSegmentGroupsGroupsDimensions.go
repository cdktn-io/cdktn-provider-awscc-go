// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition


type CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensions struct {
	// One or more calculated attributes to use as criteria for the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#calculated_attributes CustomerprofilesSegmentDefinition#calculated_attributes}
	CalculatedAttributes interface{} `field:"optional" json:"calculatedAttributes" yaml:"calculatedAttributes"`
	// Specifies the dimension settings within profile attributes for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#profile_attributes CustomerprofilesSegmentDefinition#profile_attributes}
	ProfileAttributes *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributes `field:"optional" json:"profileAttributes" yaml:"profileAttributes"`
}

