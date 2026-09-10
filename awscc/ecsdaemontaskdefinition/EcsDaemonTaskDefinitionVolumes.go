// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionVolumes struct {
	// This parameter is specified when you use bind mount host volumes.
	//
	// The contents of the ``host`` parameter determine whether your bind mount host volume persists on the host container instance and where it's stored. If the ``host`` parameter is empty, then the Docker daemon assigns a host path for your data volume. However, the data isn't guaranteed to persist after the containers that are associated with it stop running.
	//  Windows containers can mount whole directories on the same drive as ``$env:ProgramData``. Windows containers can't mount directories on a different drive, and mount point can't be across drives. For example, you can mount ``C:\my\path:C:\my\path`` and ``D:\:D:\``, but not ``D:\my\path:C:\my\path`` or ``D:\:C:\my\path``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#host EcsDaemonTaskDefinition#host}
	Host *EcsDaemonTaskDefinitionVolumesHost `field:"optional" json:"host" yaml:"host"`
	// The name of the volume.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed.
	//  When using a volume configured at launch, the ``name`` is required and must also be specified as the volume name in the ``ServiceVolumeConfiguration`` or ``TaskVolumeConfiguration`` parameter when creating your service or standalone task.
	//  For all other types of volumes, this name is referenced in the ``sourceVolume`` parameter of the ``mountPoints`` object in the container definition.
	//  When a volume is using the ``efsVolumeConfiguration``, the name is required.
	//  When a volume is using the ``s3filesVolumeConfiguration``, the name is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#name EcsDaemonTaskDefinition#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

