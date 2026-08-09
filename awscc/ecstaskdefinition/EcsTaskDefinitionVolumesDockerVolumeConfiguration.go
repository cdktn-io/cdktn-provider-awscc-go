// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionVolumesDockerVolumeConfiguration struct {
	// If this value is ``true``, the Docker volume is created if it doesn't already exist.
	//
	// This field is only used if the ``scope`` is ``shared``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_task_definition#autoprovision EcsTaskDefinition#autoprovision}
	Autoprovision interface{} `field:"optional" json:"autoprovision" yaml:"autoprovision"`
	// The Docker volume driver to use.
	//
	// The driver value must match the driver name provided by Docker because it is used for task placement. If the driver was installed using the Docker plugin CLI, use ``docker plugin ls`` to retrieve the driver name from your container instance. If the driver was installed using another method, use Docker plugin discovery to retrieve the driver name. This parameter maps to ``Driver`` in the docker container create command and the ``xxdriver`` option to docker volume create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_task_definition#driver EcsTaskDefinition#driver}
	Driver *string `field:"optional" json:"driver" yaml:"driver"`
	// A map of Docker driver-specific options passed through.
	//
	// This parameter maps to ``DriverOpts`` in the docker create-volume command and the ``xxopt`` option to docker volume create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_task_definition#driver_opts EcsTaskDefinition#driver_opts}
	DriverOpts *map[string]*string `field:"optional" json:"driverOpts" yaml:"driverOpts"`
	// Custom metadata to add to your Docker volume.
	//
	// This parameter maps to ``Labels`` in the docker container create command and the ``xxlabel`` option to docker volume create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_task_definition#labels EcsTaskDefinition#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The scope for the Docker volume that determines its lifecycle.
	//
	// Docker volumes that are scoped to a ``task`` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are scoped as ``shared`` persist after the task stops.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_task_definition#scope EcsTaskDefinition#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

