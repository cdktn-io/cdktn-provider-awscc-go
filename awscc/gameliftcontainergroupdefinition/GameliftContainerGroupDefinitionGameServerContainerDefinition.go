// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainergroupdefinition


type GameliftContainerGroupDefinitionGameServerContainerDefinition struct {
	// A descriptive label for the container definition. Container definition names must be unique with a container group definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#container_name GameliftContainerGroupDefinition#container_name}
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// A list of container dependencies that determines when this container starts up and shuts down.
	//
	// For container groups with multiple containers, dependencies let you define a startup/shutdown sequence across the containers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#depends_on GameliftContainerGroupDefinition#depends_on}
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// The environment variables to pass to a container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#environment_override GameliftContainerGroupDefinition#environment_override}
	EnvironmentOverride interface{} `field:"optional" json:"environmentOverride" yaml:"environmentOverride"`
	// Specifies the image URI of this container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#image_uri GameliftContainerGroupDefinition#image_uri}
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// Linux-specific modifications applied to the default Docker container configuration, such as Linux capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#linux_capabilities GameliftContainerGroupDefinition#linux_capabilities}
	LinuxCapabilities *GameliftContainerGroupDefinitionGameServerContainerDefinitionLinuxCapabilities `field:"optional" json:"linuxCapabilities" yaml:"linuxCapabilities"`
	// A list of mount point configurations to be used in a container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#mount_points GameliftContainerGroupDefinition#mount_points}
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// Defines the ports on the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#port_configuration GameliftContainerGroupDefinition#port_configuration}
	PortConfiguration *GameliftContainerGroupDefinitionGameServerContainerDefinitionPortConfiguration `field:"optional" json:"portConfiguration" yaml:"portConfiguration"`
	// The digest of the container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#resolved_image_digest GameliftContainerGroupDefinition#resolved_image_digest}
	ResolvedImageDigest *string `field:"optional" json:"resolvedImageDigest" yaml:"resolvedImageDigest"`
	// The version of the server SDK used in this container group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#server_sdk_version GameliftContainerGroupDefinition#server_sdk_version}
	ServerSdkVersion *string `field:"optional" json:"serverSdkVersion" yaml:"serverSdkVersion"`
}

