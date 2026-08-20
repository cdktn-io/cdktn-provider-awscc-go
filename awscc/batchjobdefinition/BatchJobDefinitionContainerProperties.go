// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionContainerProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#command BatchJobDefinition#command}.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#enable_execute_command BatchJobDefinition#enable_execute_command}.
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#environment BatchJobDefinition#environment}.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#ephemeral_storage BatchJobDefinition#ephemeral_storage}.
	EphemeralStorage *BatchJobDefinitionContainerPropertiesEphemeralStorage `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#execution_role_arn BatchJobDefinition#execution_role_arn}.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#fargate_platform_configuration BatchJobDefinition#fargate_platform_configuration}.
	FargatePlatformConfiguration *BatchJobDefinitionContainerPropertiesFargatePlatformConfiguration `field:"optional" json:"fargatePlatformConfiguration" yaml:"fargatePlatformConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#image BatchJobDefinition#image}.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#job_role_arn BatchJobDefinition#job_role_arn}.
	JobRoleArn *string `field:"optional" json:"jobRoleArn" yaml:"jobRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#linux_parameters BatchJobDefinition#linux_parameters}.
	LinuxParameters *BatchJobDefinitionContainerPropertiesLinuxParameters `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#log_configuration BatchJobDefinition#log_configuration}.
	LogConfiguration *BatchJobDefinitionContainerPropertiesLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#memory BatchJobDefinition#memory}.
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#mount_points BatchJobDefinition#mount_points}.
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#network_configuration BatchJobDefinition#network_configuration}.
	NetworkConfiguration *BatchJobDefinitionContainerPropertiesNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#privileged BatchJobDefinition#privileged}.
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#readonly_root_filesystem BatchJobDefinition#readonly_root_filesystem}.
	ReadonlyRootFilesystem interface{} `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#repository_credentials BatchJobDefinition#repository_credentials}.
	RepositoryCredentials *BatchJobDefinitionContainerPropertiesRepositoryCredentials `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#resource_requirements BatchJobDefinition#resource_requirements}.
	ResourceRequirements interface{} `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#runtime_platform BatchJobDefinition#runtime_platform}.
	RuntimePlatform *BatchJobDefinitionContainerPropertiesRuntimePlatform `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#secrets BatchJobDefinition#secrets}.
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#ulimits BatchJobDefinition#ulimits}.
	Ulimits interface{} `field:"optional" json:"ulimits" yaml:"ulimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#user BatchJobDefinition#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#vcpus BatchJobDefinition#vcpus}.
	Vcpus *float64 `field:"optional" json:"vcpus" yaml:"vcpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_job_definition#volumes BatchJobDefinition#volumes}.
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

