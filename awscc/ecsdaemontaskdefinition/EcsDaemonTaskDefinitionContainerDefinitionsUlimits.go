// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionContainerDefinitionsUlimits struct {
	// The hard limit for the ``ulimit`` type.
	//
	// The value can be specified in bytes, seconds, or as a count, depending on the ``type`` of the ``ulimit``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_daemon_task_definition#hard_limit EcsDaemonTaskDefinition#hard_limit}
	HardLimit *float64 `field:"optional" json:"hardLimit" yaml:"hardLimit"`
	// The ``type`` of the ``ulimit``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_daemon_task_definition#name EcsDaemonTaskDefinition#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The soft limit for the ``ulimit`` type.
	//
	// The value can be specified in bytes, seconds, or as a count, depending on the ``type`` of the ``ulimit``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_daemon_task_definition#soft_limit EcsDaemonTaskDefinition#soft_limit}
	SoftLimit *float64 `field:"optional" json:"softLimit" yaml:"softLimit"`
}

