// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionVolumesFsxWindowsFileServerVolumeConfiguration struct {
	// The authorization configuration details for the Amazon FSx for Windows File Server file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_definition#authorization_config EcsTaskDefinition#authorization_config}
	AuthorizationConfig *EcsTaskDefinitionVolumesFsxWindowsFileServerVolumeConfigurationAuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The Amazon FSx for Windows File Server file system ID to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_definition#file_system_id EcsTaskDefinition#file_system_id}
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// The directory within the Amazon FSx for Windows File Server file system to mount as the root directory inside the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_definition#root_directory EcsTaskDefinition#root_directory}
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
}

