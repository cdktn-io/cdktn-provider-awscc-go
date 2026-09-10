// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionContainerDefinitionsEnvironment struct {
	// The name of the key-value pair. For environment variables, this is the name of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#name EcsDaemonTaskDefinition#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the key-value pair. For environment variables, this is the value of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#value EcsDaemonTaskDefinition#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

