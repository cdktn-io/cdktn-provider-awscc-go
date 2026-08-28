// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainergroupdefinition


type GameliftContainerGroupDefinitionSupportContainerDefinitions struct {
	// A descriptive label for the container definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#container_name GameliftContainerGroupDefinition#container_name}
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// A list of container dependencies that determines when this container starts up and shuts down.
	//
	// For container groups with multiple containers, dependencies let you define a startup/shutdown sequence across the containers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#depends_on GameliftContainerGroupDefinition#depends_on}
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// The environment variables to pass to a container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#environment_override GameliftContainerGroupDefinition#environment_override}
	EnvironmentOverride interface{} `field:"optional" json:"environmentOverride" yaml:"environmentOverride"`
	// Specifies if the container is essential.
	//
	// If an essential container fails a health check, then all containers in the container group will be restarted. You must specify exactly 1 essential container in a container group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#essential GameliftContainerGroupDefinition#essential}
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// Specifies how the health of the containers will be checked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#health_check GameliftContainerGroupDefinition#health_check}
	HealthCheck *GameliftContainerGroupDefinitionSupportContainerDefinitionsHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Specifies the image URI of this container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#image_uri GameliftContainerGroupDefinition#image_uri}
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// Linux-specific modifications applied to the default Docker container configuration, such as Linux capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#linux_capabilities GameliftContainerGroupDefinition#linux_capabilities}
	LinuxCapabilities *GameliftContainerGroupDefinitionSupportContainerDefinitionsLinuxCapabilities `field:"optional" json:"linuxCapabilities" yaml:"linuxCapabilities"`
	// The total memory limit of container groups following this definition in MiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#memory_hard_limit_mebibytes GameliftContainerGroupDefinition#memory_hard_limit_mebibytes}
	MemoryHardLimitMebibytes *float64 `field:"optional" json:"memoryHardLimitMebibytes" yaml:"memoryHardLimitMebibytes"`
	// A list of mount point configurations to be used in a container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#mount_points GameliftContainerGroupDefinition#mount_points}
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// Defines the ports on the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#port_configuration GameliftContainerGroupDefinition#port_configuration}
	PortConfiguration *GameliftContainerGroupDefinitionSupportContainerDefinitionsPortConfiguration `field:"optional" json:"portConfiguration" yaml:"portConfiguration"`
	// The digest of the container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#resolved_image_digest GameliftContainerGroupDefinition#resolved_image_digest}
	ResolvedImageDigest *string `field:"optional" json:"resolvedImageDigest" yaml:"resolvedImageDigest"`
	// The number of virtual CPUs to give to the support group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/gamelift_container_group_definition#vcpu GameliftContainerGroupDefinition#vcpu}
	Vcpu *float64 `field:"optional" json:"vcpu" yaml:"vcpu"`
}

