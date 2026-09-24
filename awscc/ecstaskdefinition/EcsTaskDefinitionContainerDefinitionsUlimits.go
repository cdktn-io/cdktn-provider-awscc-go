// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsUlimits struct {
	// The hard limit for the ``ulimit`` type.
	//
	// The value can be specified in bytes, seconds, or as a count, depending on the ``type`` of the ``ulimit``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_task_definition#hard_limit EcsTaskDefinition#hard_limit}
	HardLimit *float64 `field:"optional" json:"hardLimit" yaml:"hardLimit"`
	// The ``type`` of the ``ulimit``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_task_definition#name EcsTaskDefinition#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The soft limit for the ``ulimit`` type.
	//
	// The value can be specified in bytes, seconds, or as a count, depending on the ``type`` of the ``ulimit``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_task_definition#soft_limit EcsTaskDefinition#soft_limit}
	SoftLimit *float64 `field:"optional" json:"softLimit" yaml:"softLimit"`
}

