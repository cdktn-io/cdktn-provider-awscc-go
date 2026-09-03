// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionContainerDefinitionsLinuxParameters struct {
	// The Linux capabilities for the container that are added to or dropped from the default configuration provided by Docker.
	//
	// For tasks that use the Fargate launch type, ``capabilities`` is supported for all platform versions but the ``add`` parameter is only supported if using platform version 1.4.0 or later.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecs_daemon_task_definition#capabilities EcsDaemonTaskDefinition#capabilities}
	Capabilities *EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersCapabilities `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Any host devices to expose to the container.
	//
	// This parameter maps to ``Devices`` in the docker container create command and the ``--device`` option to docker run.
	//   If you're using tasks that use the Fargate launch type, the ``devices`` parameter isn't supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecs_daemon_task_definition#devices EcsDaemonTaskDefinition#devices}
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// Run an ``init`` process inside the container that forwards signals and reaps processes.
	//
	// This parameter maps to the ``--init`` option to docker run. This parameter requires version 1.25 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: ``sudo docker version --format '{{.Server.APIVersion}}'``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecs_daemon_task_definition#init_process_enabled EcsDaemonTaskDefinition#init_process_enabled}
	InitProcessEnabled interface{} `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// The container path, mount options, and size (in MiB) of the tmpfs mount.
	//
	// This parameter maps to the ``--tmpfs`` option to docker run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ecs_daemon_task_definition#tmpfs EcsDaemonTaskDefinition#tmpfs}
	Tmpfs interface{} `field:"optional" json:"tmpfs" yaml:"tmpfs"`
}

