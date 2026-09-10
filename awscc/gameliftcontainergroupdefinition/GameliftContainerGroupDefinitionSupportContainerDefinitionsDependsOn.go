// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainergroupdefinition


type GameliftContainerGroupDefinitionSupportContainerDefinitionsDependsOn struct {
	// The type of dependency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/gamelift_container_group_definition#condition GameliftContainerGroupDefinition#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// A descriptive label for the container definition. The container being defined depends on this container's condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/gamelift_container_group_definition#container_name GameliftContainerGroupDefinition#container_name}
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
}

