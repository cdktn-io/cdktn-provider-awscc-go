// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionVolumesHost struct {
	// When the ``host`` parameter is used, specify a ``sourcePath`` to declare the path on the host container instance that's presented to the container.
	//
	// If this parameter is empty, then the Docker daemon has assigned a host path for you. If the ``host`` parameter contains a ``sourcePath`` file location, then the data volume persists at the specified location on the host container instance until you delete it manually. If the ``sourcePath`` value doesn't exist on the host container instance, the Docker daemon creates it. If the location does exist, the contents of the source path folder are exported.
	//  If you're using the Fargate launch type, the ``sourcePath`` parameter is not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_daemon_task_definition#source_path EcsDaemonTaskDefinition#source_path}
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

