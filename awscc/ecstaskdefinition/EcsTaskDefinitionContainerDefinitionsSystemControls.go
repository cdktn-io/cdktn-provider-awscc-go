// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsSystemControls struct {
	// The namespaced kernel parameter to set a ``value`` for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecs_task_definition#namespace EcsTaskDefinition#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The namespaced kernel parameter to set a ``value`` for.
	//
	// Valid IPC namespace values: ``"kernel.msgmax" | "kernel.msgmnb" | "kernel.msgmni" | "kernel.sem" | "kernel.shmall" | "kernel.shmmax" | "kernel.shmmni" | "kernel.shm_rmid_forced"``, and ``Sysctls`` that start with ``"fs.mqueue.*"``
	//  Valid network namespace values: ``Sysctls`` that start with ``"net.*"``. Only namespaced ``Sysctls`` that exist within the container starting with "net.* are accepted.
	//  All of these values are supported by Fargate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecs_task_definition#value EcsTaskDefinition#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

