// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtask

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsmMaintenanceWindowTaskConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The priority of the task in the maintenance window.
	//
	// The lower the number, the higher the priority. Tasks that have the same priority are scheduled in parallel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#priority SsmMaintenanceWindowTask#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The resource that the task uses during execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#task_arn SsmMaintenanceWindowTask#task_arn}
	TaskArn *string `field:"required" json:"taskArn" yaml:"taskArn"`
	// The type of task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#task_type SsmMaintenanceWindowTask#task_type}
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// The ID of the maintenance window where the task is registered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#window_id SsmMaintenanceWindowTask#window_id}
	WindowId *string `field:"required" json:"windowId" yaml:"windowId"`
	// The specification for whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#cutoff_behavior SsmMaintenanceWindowTask#cutoff_behavior}
	CutoffBehavior *string `field:"optional" json:"cutoffBehavior" yaml:"cutoffBehavior"`
	// A description of the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#description SsmMaintenanceWindowTask#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Information about an Amazon S3 bucket to write Run Command task-level logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#logging_info SsmMaintenanceWindowTask#logging_info}
	LoggingInfo *SsmMaintenanceWindowTaskLoggingInfo `field:"optional" json:"loggingInfo" yaml:"loggingInfo"`
	// The maximum number of targets this task can be run for, in parallel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#max_concurrency SsmMaintenanceWindowTask#max_concurrency}
	MaxConcurrency *string `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The maximum number of errors allowed before this task stops being scheduled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#max_errors SsmMaintenanceWindowTask#max_errors}
	MaxErrors *string `field:"optional" json:"maxErrors" yaml:"maxErrors"`
	// The task name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#name SsmMaintenanceWindowTask#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the IAM service role for AWS Systems Manager to assume when running a maintenance window task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#service_role_arn SsmMaintenanceWindowTask#service_role_arn}
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// The targets (either instances or window target ids).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#targets SsmMaintenanceWindowTask#targets}
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// The parameters to pass to the task when it runs.
	//
	// Populate only the fields that match the task type. All other fields should be empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#task_invocation_parameters SsmMaintenanceWindowTask#task_invocation_parameters}
	TaskInvocationParameters *SsmMaintenanceWindowTaskTaskInvocationParameters `field:"optional" json:"taskInvocationParameters" yaml:"taskInvocationParameters"`
	// The parameters to pass to the task when it runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window_task#task_parameters SsmMaintenanceWindowTask#task_parameters}
	TaskParameters *string `field:"optional" json:"taskParameters" yaml:"taskParameters"`
}

