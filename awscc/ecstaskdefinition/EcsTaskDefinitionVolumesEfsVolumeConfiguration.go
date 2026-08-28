// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionVolumesEfsVolumeConfiguration struct {
	// The authorization configuration details for the Amazon EFS file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_definition#authorization_config EcsTaskDefinition#authorization_config}
	AuthorizationConfig *EcsTaskDefinitionVolumesEfsVolumeConfigurationAuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The Amazon EFS file system ID to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_definition#filesystem_id EcsTaskDefinition#filesystem_id}
	FilesystemId *string `field:"optional" json:"filesystemId" yaml:"filesystemId"`
	// The directory within the Amazon EFS file system to mount as the root directory inside the host.
	//
	// If this parameter is omitted, the root of the Amazon EFS volume will be used. Specifying ``/`` will have the same effect as omitting this parameter.
	//   If an EFS access point is specified in the ``authorizationConfig``, the root directory parameter must either be omitted or set to ``/`` which will enforce the path set on the EFS access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_definition#root_directory EcsTaskDefinition#root_directory}
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// Determines whether to use encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server.
	//
	// Transit encryption must be turned on if Amazon EFS IAM authorization is used. If this parameter is omitted, the default value of ``DISABLED`` is used. For more information, see [Encrypting data in transit](https://docs.aws.amazon.com/efs/latest/ug/encryption-in-transit.html) in the *Amazon Elastic File System User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_definition#transit_encryption EcsTaskDefinition#transit_encryption}
	TransitEncryption *string `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
	// The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server.
	//
	// If you do not specify a transit encryption port, it will use the port selection strategy that the Amazon EFS mount helper uses. For more information, see [EFS mount helper](https://docs.aws.amazon.com/efs/latest/ug/efs-mount-helper.html) in the *Amazon Elastic File System User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_definition#transit_encryption_port EcsTaskDefinition#transit_encryption_port}
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}

