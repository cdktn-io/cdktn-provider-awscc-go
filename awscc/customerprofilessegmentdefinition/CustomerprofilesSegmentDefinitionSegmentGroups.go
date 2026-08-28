// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition


type CustomerprofilesSegmentDefinitionSegmentGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_segment_definition#groups CustomerprofilesSegmentDefinition#groups}.
	Groups interface{} `field:"optional" json:"groups" yaml:"groups"`
	// Specifies the operator on how to handle multiple groups within the same segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_segment_definition#include CustomerprofilesSegmentDefinition#include}
	Include *string `field:"optional" json:"include" yaml:"include"`
}

