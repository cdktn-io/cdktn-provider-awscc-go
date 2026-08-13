// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition


type CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributes struct {
	// Overrides the condition block within the original calculated attribute definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/customerprofiles_segment_definition#condition_overrides CustomerprofilesSegmentDefinition#condition_overrides}
	ConditionOverrides *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsCalculatedAttributesConditionOverrides `field:"optional" json:"conditionOverrides" yaml:"conditionOverrides"`
	// The type of segment dimension to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/customerprofiles_segment_definition#dimension_type CustomerprofilesSegmentDefinition#dimension_type}
	DimensionType *string `field:"optional" json:"dimensionType" yaml:"dimensionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/customerprofiles_segment_definition#values CustomerprofilesSegmentDefinition#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

