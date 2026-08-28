// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersDevices struct {
	// The path inside the container at which to expose the host device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_daemon_task_definition#container_path EcsDaemonTaskDefinition#container_path}
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The path for the device on the host container instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_daemon_task_definition#host_path EcsDaemonTaskDefinition#host_path}
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
	// The explicit permissions to provide to the container for the device.
	//
	// By default, the container has permissions for ``read``, ``write``, and ``mknod`` for the device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_daemon_task_definition#permissions EcsDaemonTaskDefinition#permissions}
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
}

