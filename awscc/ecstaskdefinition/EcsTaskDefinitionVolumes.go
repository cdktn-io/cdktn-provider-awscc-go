// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionVolumes struct {
	// Indicates whether the volume should be configured at launch time.
	//
	// This is used to create Amazon EBS volumes for standalone tasks or tasks created as part of a service. Each task definition revision may only have one volume configured at launch in the volume configuration.
	//  To configure a volume at launch time, use this task definition revision and specify a ``volumeConfigurations`` object when calling the ``CreateService``, ``UpdateService``, ``RunTask`` or ``StartTask`` APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_definition#configured_at_launch EcsTaskDefinition#configured_at_launch}
	ConfiguredAtLaunch interface{} `field:"optional" json:"configuredAtLaunch" yaml:"configuredAtLaunch"`
	// This parameter is specified when you use Docker volumes.
	//
	// Windows containers only support the use of the ``local`` driver. To use bind mounts, specify the ``host`` parameter instead.
	//   Docker volumes aren't supported by tasks run on FARGATElong.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_definition#docker_volume_configuration EcsTaskDefinition#docker_volume_configuration}
	DockerVolumeConfiguration *EcsTaskDefinitionVolumesDockerVolumeConfiguration `field:"optional" json:"dockerVolumeConfiguration" yaml:"dockerVolumeConfiguration"`
	// This parameter is specified when you use an Amazon Elastic File System file system for task storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_definition#efs_volume_configuration EcsTaskDefinition#efs_volume_configuration}
	EfsVolumeConfiguration *EcsTaskDefinitionVolumesEfsVolumeConfiguration `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// This parameter is specified when you use Amazon FSx for Windows File Server file system for task storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_definition#fsx_windows_file_server_volume_configuration EcsTaskDefinition#fsx_windows_file_server_volume_configuration}
	FsxWindowsFileServerVolumeConfiguration *EcsTaskDefinitionVolumesFsxWindowsFileServerVolumeConfiguration `field:"optional" json:"fsxWindowsFileServerVolumeConfiguration" yaml:"fsxWindowsFileServerVolumeConfiguration"`
	// This parameter is specified when you use bind mount host volumes.
	//
	// The contents of the ``host`` parameter determine whether your bind mount host volume persists on the host container instance and where it's stored. If the ``host`` parameter is empty, then the Docker daemon assigns a host path for your data volume. However, the data isn't guaranteed to persist after the containers that are associated with it stop running.
	//  Windows containers can mount whole directories on the same drive as ``$env:ProgramData``. Windows containers can't mount directories on a different drive, and mount point can't be across drives. For example, you can mount ``C:\my\path:C:\my\path`` and ``D:\:D:\``, but not ``D:\my\path:C:\my\path`` or ``D:\:C:\my\path``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_definition#host EcsTaskDefinition#host}
	Host *EcsTaskDefinitionVolumesHost `field:"optional" json:"host" yaml:"host"`
	// The name of the volume.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed.
	//  When using a volume configured at launch, the ``name`` is required and must also be specified as the volume name in the ``ServiceVolumeConfiguration`` or ``TaskVolumeConfiguration`` parameter when creating your service or standalone task.
	//  For all other types of volumes, this name is referenced in the ``sourceVolume`` parameter of the ``mountPoints`` object in the container definition.
	//  When a volume is using the ``efsVolumeConfiguration``, the name is required.
	//  When a volume is using the ``s3filesVolumeConfiguration``, the name is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_definition#name EcsTaskDefinition#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// This parameter is specified when you use an Amazon S3 Files file system for task storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_task_definition#s3_files_volume_configuration EcsTaskDefinition#s3_files_volume_configuration}
	S3FilesVolumeConfiguration *EcsTaskDefinitionVolumesS3FilesVolumeConfiguration `field:"optional" json:"s3FilesVolumeConfiguration" yaml:"s3FilesVolumeConfiguration"`
}

