// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisetask


type IotsitewiseTaskTaskConfigurationContainerTaskConfiguration struct {
	// The Amazon ECR image URI for the task container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_task#ecr_uri IotsitewiseTask#ecr_uri}
	EcrUri *string `field:"required" json:"ecrUri" yaml:"ecrUri"`
	// The processing type for compute resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_task#processing_type IotsitewiseTask#processing_type}
	ProcessingType *string `field:"required" json:"processingType" yaml:"processingType"`
	// The processing unit allocation that determines vCPU, memory, and GPU resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_task#processing_unit IotsitewiseTask#processing_unit}
	ProcessingUnit *string `field:"required" json:"processingUnit" yaml:"processingUnit"`
	// The ARN of the IAM role that grants the containerized workload permissions to access AWS resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_task#task_execution_role IotsitewiseTask#task_execution_role}
	TaskExecutionRole *string `field:"required" json:"taskExecutionRole" yaml:"taskExecutionRole"`
	// The command to execute in the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_task#command IotsitewiseTask#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// A map of environment variable key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_task#environment_variables IotsitewiseTask#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The timeout in seconds for task execution. Default: 3600 (1 hour).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_task#timeout_seconds IotsitewiseTask#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

