// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition


type CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesConditionOverridesRange struct {
	// The ending point for this overridden range.
	//
	// Positive numbers indicate how many days in the past data should be included, and negative numbers indicate how many days in the future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#end CustomerprofilesSegmentDefinition#end}
	End *float64 `field:"optional" json:"end" yaml:"end"`
	// The starting point for this overridden range.
	//
	// Positive numbers indicate how many days in the past data should be included, and negative numbers indicate how many days in the future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#start CustomerprofilesSegmentDefinition#start}
	Start *float64 `field:"optional" json:"start" yaml:"start"`
	// The unit to be applied to the range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#unit CustomerprofilesSegmentDefinition#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

