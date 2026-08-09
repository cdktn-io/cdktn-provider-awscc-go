// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainergroupdefinition

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GameliftContainerGroupDefinitionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A descriptive label for the container group definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#name GameliftContainerGroupDefinition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The operating system of the container group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#operating_system GameliftContainerGroupDefinition#operating_system}
	OperatingSystem *string `field:"required" json:"operatingSystem" yaml:"operatingSystem"`
	// The total memory limit of container groups following this definition in MiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#total_memory_limit_mebibytes GameliftContainerGroupDefinition#total_memory_limit_mebibytes}
	TotalMemoryLimitMebibytes *float64 `field:"required" json:"totalMemoryLimitMebibytes" yaml:"totalMemoryLimitMebibytes"`
	// The total amount of virtual CPUs on the container group definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#total_vcpu_limit GameliftContainerGroupDefinition#total_vcpu_limit}
	TotalVcpuLimit *float64 `field:"required" json:"totalVcpuLimit" yaml:"totalVcpuLimit"`
	// The scope of the container group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#container_group_type GameliftContainerGroupDefinition#container_group_type}
	ContainerGroupType *string `field:"optional" json:"containerGroupType" yaml:"containerGroupType"`
	// Specifies the information required to run game servers with this container group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#game_server_container_definition GameliftContainerGroupDefinition#game_server_container_definition}
	GameServerContainerDefinition *GameliftContainerGroupDefinitionGameServerContainerDefinition `field:"optional" json:"gameServerContainerDefinition" yaml:"gameServerContainerDefinition"`
	// A specific ContainerGroupDefinition version to be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#source_version_number GameliftContainerGroupDefinition#source_version_number}
	SourceVersionNumber *float64 `field:"optional" json:"sourceVersionNumber" yaml:"sourceVersionNumber"`
	// A collection of support container definitions that define the containers in this group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#support_container_definitions GameliftContainerGroupDefinition#support_container_definitions}
	SupportContainerDefinitions interface{} `field:"optional" json:"supportContainerDefinitions" yaml:"supportContainerDefinitions"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#tags GameliftContainerGroupDefinition#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The description of this version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/gamelift_container_group_definition#version_description GameliftContainerGroupDefinition#version_description}
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

