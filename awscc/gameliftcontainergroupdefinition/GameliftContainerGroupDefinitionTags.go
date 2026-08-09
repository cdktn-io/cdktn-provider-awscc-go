// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainergroupdefinition


type GameliftContainerGroupDefinitionTags struct {
	// The key name of the tag.
	//
	// You can specify a value that is 1 to 128 Unicode characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#key GameliftContainerGroupDefinition#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#value GameliftContainerGroupDefinition#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

