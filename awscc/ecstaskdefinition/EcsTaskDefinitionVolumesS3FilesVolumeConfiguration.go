// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionVolumesS3FilesVolumeConfiguration struct {
	// The full ARN of the S3 Files access point to use.
	//
	// If an access point is specified, the root directory value specified in the ``S3FilesVolumeConfiguration`` must either be omitted or set to ``/`` which will enforce the path set on the S3 Files access point. For more information, see [Creating S3 Files access points](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-files-access-points-creating.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_task_definition#access_point_arn EcsTaskDefinition#access_point_arn}
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// The full ARN of the S3 Files file system to mount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_task_definition#file_system_arn EcsTaskDefinition#file_system_arn}
	FileSystemArn *string `field:"optional" json:"fileSystemArn" yaml:"fileSystemArn"`
	// The directory within the Amazon S3 Files file system to mount as the root directory.
	//
	// If this parameter is omitted, the root of the Amazon S3 Files file system will be used. Specifying ``/`` will have the same effect as omitting this parameter.
	//   If a S3 Files access point is specified in the ``accessPointArn``, the root directory parameter must either be omitted or set to ``/`` which will enforce the path set on the S3 Files access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_task_definition#root_directory EcsTaskDefinition#root_directory}
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// The port to use for sending encrypted data between the ECS host and the S3 Files file system.
	//
	// If you do not specify a transit encryption port, it will use the port selection strategy that the Amazon S3 Files mount helper uses. For more information, see [S3 Files mount helper](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-files-mounting.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_task_definition#transit_encryption_port EcsTaskDefinition#transit_encryption_port}
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}

