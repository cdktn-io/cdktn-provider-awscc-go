// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainergroupdefinition


type GameliftContainerGroupDefinitionGameServerContainerDefinitionMountPoints struct {
	// The access permissions for the mounted path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#access_level GameliftContainerGroupDefinition#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// The path inside the container where the mount is accessible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#container_path GameliftContainerGroupDefinition#container_path}
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The path on the host that will be mounted in the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#instance_path GameliftContainerGroupDefinition#instance_path}
	InstancePath *string `field:"optional" json:"instancePath" yaml:"instancePath"`
}

