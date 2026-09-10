// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionContainerDefinitions struct {
	// The command that's passed to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#command EcsDaemonTaskDefinition#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The number of ``cpu`` units reserved for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#cpu EcsDaemonTaskDefinition#cpu}
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The dependencies defined for container startup and shutdown.
	//
	// A container can contain multiple dependencies on other containers in a task definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#depends_on EcsDaemonTaskDefinition#depends_on}
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// The entry point that's passed to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#entry_point EcsDaemonTaskDefinition#entry_point}
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The environment variables to pass to a container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#environment EcsDaemonTaskDefinition#environment}
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// A list of files containing the environment variables to pass to a container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#environment_files EcsDaemonTaskDefinition#environment_files}
	EnvironmentFiles interface{} `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// If the ``essential`` parameter of a container is marked as ``true``, and that container fails or stops for any reason, all other containers that are part of the task are stopped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#essential EcsDaemonTaskDefinition#essential}
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// The FireLens configuration for the container. This is used to specify and configure a log router for container logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#firelens_configuration EcsDaemonTaskDefinition#firelens_configuration}
	FirelensConfiguration *EcsDaemonTaskDefinitionContainerDefinitionsFirelensConfiguration `field:"optional" json:"firelensConfiguration" yaml:"firelensConfiguration"`
	// The container health check command and associated configuration parameters for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#health_check EcsDaemonTaskDefinition#health_check}
	HealthCheck *EcsDaemonTaskDefinitionContainerDefinitionsHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The image used to start the container.
	//
	// This string is passed directly to the Docker daemon. Images in the Docker Hub registry are available by default. Other repositories are specified with either ``repository-url/image:tag`` or ``repository-url/image@digest``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#image EcsDaemonTaskDefinition#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// When this parameter is ``true``, you can deploy containerized applications that require ``stdin`` or a ``tty`` to be allocated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#interactive EcsDaemonTaskDefinition#interactive}
	Interactive interface{} `field:"optional" json:"interactive" yaml:"interactive"`
	// Linux-specific modifications that are applied to the container configuration, such as Linux kernel capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#linux_parameters EcsDaemonTaskDefinition#linux_parameters}
	LinuxParameters *EcsDaemonTaskDefinitionContainerDefinitionsLinuxParameters `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// The log configuration specification for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#log_configuration EcsDaemonTaskDefinition#log_configuration}
	LogConfiguration *EcsDaemonTaskDefinitionContainerDefinitionsLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If the container attempts to exceed the memory specified here, the container is killed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#memory EcsDaemonTaskDefinition#memory}
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#memory_reservation EcsDaemonTaskDefinition#memory_reservation}
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// The mount points for data volumes in your container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#mount_points EcsDaemonTaskDefinition#mount_points}
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// The name of the container. Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#name EcsDaemonTaskDefinition#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the ``root`` user).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#privileged EcsDaemonTaskDefinition#privileged}
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// When this parameter is ``true``, a TTY is allocated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#pseudo_terminal EcsDaemonTaskDefinition#pseudo_terminal}
	PseudoTerminal interface{} `field:"optional" json:"pseudoTerminal" yaml:"pseudoTerminal"`
	// When this parameter is true, the container is given read-only access to its root file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#readonly_root_filesystem EcsDaemonTaskDefinition#readonly_root_filesystem}
	ReadonlyRootFilesystem interface{} `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// The private repository authentication credentials to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#repository_credentials EcsDaemonTaskDefinition#repository_credentials}
	RepositoryCredentials *EcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentials `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
	// The restart policy for the container.
	//
	// When you set up a restart policy, Amazon ECS can restart the container without needing to replace the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#restart_policy EcsDaemonTaskDefinition#restart_policy}
	RestartPolicy *EcsDaemonTaskDefinitionContainerDefinitionsRestartPolicy `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// The secrets to pass to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#secrets EcsDaemonTaskDefinition#secrets}
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
	// Time duration (in seconds) to wait before giving up on resolving dependencies for a container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#start_timeout EcsDaemonTaskDefinition#start_timeout}
	StartTimeout *float64 `field:"optional" json:"startTimeout" yaml:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#stop_timeout EcsDaemonTaskDefinition#stop_timeout}
	StopTimeout *float64 `field:"optional" json:"stopTimeout" yaml:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#system_controls EcsDaemonTaskDefinition#system_controls}
	SystemControls interface{} `field:"optional" json:"systemControls" yaml:"systemControls"`
	// A list of ``ulimits`` to set in the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#ulimits EcsDaemonTaskDefinition#ulimits}
	Ulimits interface{} `field:"optional" json:"ulimits" yaml:"ulimits"`
	// The user to use inside the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#user EcsDaemonTaskDefinition#user}
	User *string `field:"optional" json:"user" yaml:"user"`
	// The working directory to run commands inside the container in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_daemon_task_definition#working_directory EcsDaemonTaskDefinition#working_directory}
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

