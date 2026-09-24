// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsLinuxParameters struct {
	// The Linux capabilities for the container that are added to or dropped from the default configuration provided by Docker.
	//
	// For tasks that use the Fargate launch type, ``capabilities`` is supported for all platform versions but the ``add`` parameter is only supported if using platform version 1.4.0 or later.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_task_definition#capabilities EcsTaskDefinition#capabilities}
	Capabilities *EcsTaskDefinitionContainerDefinitionsLinuxParametersCapabilities `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Any host devices to expose to the container.
	//
	// This parameter maps to ``Devices`` in the docker container create command and the ``--device`` option to docker run.
	//   If you're using tasks that use the Fargate launch type, the ``devices`` parameter isn't supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_task_definition#devices EcsTaskDefinition#devices}
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// Run an ``init`` process inside the container that forwards signals and reaps processes.
	//
	// This parameter maps to the ``--init`` option to docker run. This parameter requires version 1.25 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: ``sudo docker version --format '{{.Server.APIVersion}}'``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_task_definition#init_process_enabled EcsTaskDefinition#init_process_enabled}
	InitProcessEnabled interface{} `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// The total amount of swap memory (in MiB) a container can use.
	//
	// This parameter will be translated to the ``--memory-swap`` option to docker run where the value would be the sum of the container memory plus the ``maxSwap`` value.
	//  If a ``maxSwap`` value of ``0`` is specified, the container will not use swap. Accepted values are ``0`` or any positive integer. If the ``maxSwap`` parameter is omitted, the container will use the swap configuration for the container instance it is running on. A ``maxSwap`` value must be set for the ``swappiness`` parameter to be used.
	//   If you're using tasks that use the Fargate launch type, the ``maxSwap`` parameter isn't supported.
	//  If you're using tasks on Amazon Linux 2023 the ``swappiness`` parameter isn't supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_task_definition#max_swap EcsTaskDefinition#max_swap}
	MaxSwap *float64 `field:"optional" json:"maxSwap" yaml:"maxSwap"`
	// The value for the size (in MiB) of the ``/dev/shm`` volume.
	//
	// This parameter maps to the ``--shm-size`` option to docker run.
	//   If you are using tasks that use the Fargate launch type, the ``sharedMemorySize`` parameter is not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_task_definition#shared_memory_size EcsTaskDefinition#shared_memory_size}
	SharedMemorySize *float64 `field:"optional" json:"sharedMemorySize" yaml:"sharedMemorySize"`
	// This allows you to tune a container's memory swappiness behavior.
	//
	// A ``swappiness`` value of ``0`` will cause swapping to not happen unless absolutely necessary. A ``swappiness`` value of ``100`` will cause pages to be swapped very aggressively. Accepted values are whole numbers between ``0`` and ``100``. If the ``swappiness`` parameter is not specified, a default value of ``60`` is used. If a value is not specified for ``maxSwap`` then this parameter is ignored. This parameter maps to the ``--memory-swappiness`` option to docker run.
	//   If you're using tasks that use the Fargate launch type, the ``swappiness`` parameter isn't supported.
	//  If you're using tasks on Amazon Linux 2023 the ``swappiness`` parameter isn't supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_task_definition#swappiness EcsTaskDefinition#swappiness}
	Swappiness *float64 `field:"optional" json:"swappiness" yaml:"swappiness"`
	// The container path, mount options, and size (in MiB) of the tmpfs mount.
	//
	// This parameter maps to the ``--tmpfs`` option to docker run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_task_definition#tmpfs EcsTaskDefinition#tmpfs}
	Tmpfs interface{} `field:"optional" json:"tmpfs" yaml:"tmpfs"`
}

