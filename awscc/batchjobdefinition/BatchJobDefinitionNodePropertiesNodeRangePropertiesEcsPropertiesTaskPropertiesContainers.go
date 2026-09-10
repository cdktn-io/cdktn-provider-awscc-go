// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesContainers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#command BatchJobDefinition#command}.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#depends_on BatchJobDefinition#depends_on}.
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#environment BatchJobDefinition#environment}.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#essential BatchJobDefinition#essential}.
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#firelens_configuration BatchJobDefinition#firelens_configuration}.
	FirelensConfiguration *BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesContainersFirelensConfiguration `field:"optional" json:"firelensConfiguration" yaml:"firelensConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#image BatchJobDefinition#image}.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#linux_parameters BatchJobDefinition#linux_parameters}.
	LinuxParameters *BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesContainersLinuxParameters `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#log_configuration BatchJobDefinition#log_configuration}.
	LogConfiguration *BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesContainersLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#mount_points BatchJobDefinition#mount_points}.
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#name BatchJobDefinition#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#privileged BatchJobDefinition#privileged}.
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#readonly_root_filesystem BatchJobDefinition#readonly_root_filesystem}.
	ReadonlyRootFilesystem interface{} `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#repository_credentials BatchJobDefinition#repository_credentials}.
	RepositoryCredentials *BatchJobDefinitionNodePropertiesNodeRangePropertiesEcsPropertiesTaskPropertiesContainersRepositoryCredentials `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#resource_requirements BatchJobDefinition#resource_requirements}.
	ResourceRequirements interface{} `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#secrets BatchJobDefinition#secrets}.
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#start_timeout BatchJobDefinition#start_timeout}.
	StartTimeout *float64 `field:"optional" json:"startTimeout" yaml:"startTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#stop_timeout BatchJobDefinition#stop_timeout}.
	StopTimeout *float64 `field:"optional" json:"stopTimeout" yaml:"stopTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#ulimits BatchJobDefinition#ulimits}.
	Ulimits interface{} `field:"optional" json:"ulimits" yaml:"ulimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#user BatchJobDefinition#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
}

