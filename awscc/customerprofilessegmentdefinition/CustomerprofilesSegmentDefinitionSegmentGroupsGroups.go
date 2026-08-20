// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition


type CustomerprofilesSegmentDefinitionSegmentGroupsGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_segment_definition#dimensions CustomerprofilesSegmentDefinition#dimensions}.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_segment_definition#source_segments CustomerprofilesSegmentDefinition#source_segments}.
	SourceSegments interface{} `field:"optional" json:"sourceSegments" yaml:"sourceSegments"`
	// Specifies the operator on how to handle multiple groups within the same segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_segment_definition#source_type CustomerprofilesSegmentDefinition#source_type}
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// Specifies the operator on how to handle multiple groups within the same segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_segment_definition#type CustomerprofilesSegmentDefinition#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

