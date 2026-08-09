// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionProxyConfigurationProxyConfigurationProperties struct {
	// The name of the key-value pair. For environment variables, this is the name of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_task_definition#name EcsTaskDefinition#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the key-value pair. For environment variables, this is the value of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_task_definition#value EcsTaskDefinition#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

