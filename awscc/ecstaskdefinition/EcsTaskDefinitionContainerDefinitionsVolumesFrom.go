// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsVolumesFrom struct {
	// If this value is ``true``, the container has read-only access to the volume.
	//
	// If this value is ``false``, then the container can write to the volume. The default value is ``false``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_task_definition#read_only EcsTaskDefinition#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The name of another container within the same task definition to mount volumes from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ecs_task_definition#source_container EcsTaskDefinition#source_container}
	SourceContainer *string `field:"optional" json:"sourceContainer" yaml:"sourceContainer"`
}

