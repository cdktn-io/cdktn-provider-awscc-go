// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsMountPoints struct {
	// The path on the container to mount the host volume at.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_definition#container_path EcsTaskDefinition#container_path}
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// If this value is ``true``, the container has read-only access to the volume.
	//
	// If this value is ``false``, then the container can write to the volume. The default value is ``false``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_definition#read_only EcsTaskDefinition#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The name of the volume to mount.
	//
	// Must be a volume name referenced in the ``name`` parameter of task definition ``volume``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_task_definition#source_volume EcsTaskDefinition#source_volume}
	SourceVolume *string `field:"optional" json:"sourceVolume" yaml:"sourceVolume"`
}

