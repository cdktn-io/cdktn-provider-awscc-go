// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluejob

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueJobConfig struct {
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
	// The code that executes a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#command GlueJob#command}
	Command *GlueJobCommand `field:"required" json:"command" yaml:"command"`
	// The name or Amazon Resource Name (ARN) of the IAM role associated with this job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#role GlueJob#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The number of capacity units that are allocated to this job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#allocated_capacity GlueJob#allocated_capacity}
	AllocatedCapacity *float64 `field:"optional" json:"allocatedCapacity" yaml:"allocatedCapacity"`
	// Specifies the connections used by a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#connections GlueJob#connections}
	Connections *GlueJobConnections `field:"optional" json:"connections" yaml:"connections"`
	// The default arguments for this job, specified as name-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#default_arguments GlueJob#default_arguments}
	DefaultArguments *string `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	// A description of the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#description GlueJob#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether the job is run with a standard or flexible execution class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#execution_class GlueJob#execution_class}
	ExecutionClass *string `field:"optional" json:"executionClass" yaml:"executionClass"`
	// The maximum number of concurrent runs that are allowed for this job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#execution_property GlueJob#execution_property}
	ExecutionProperty *GlueJobExecutionProperty `field:"optional" json:"executionProperty" yaml:"executionProperty"`
	// Glue version determines the versions of Apache Spark and Python that AWS Glue supports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#glue_version GlueJob#glue_version}
	GlueVersion *string `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// Property description not available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#job_mode GlueJob#job_mode}
	JobMode *string `field:"optional" json:"jobMode" yaml:"jobMode"`
	// Property description not available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#job_run_queuing_enabled GlueJob#job_run_queuing_enabled}
	JobRunQueuingEnabled interface{} `field:"optional" json:"jobRunQueuingEnabled" yaml:"jobRunQueuingEnabled"`
	// This field is reserved for future use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#log_uri GlueJob#log_uri}
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// Property description not available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#maintenance_window GlueJob#maintenance_window}
	MaintenanceWindow *string `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#max_capacity GlueJob#max_capacity}
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum number of times to retry this job after a JobRun fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#max_retries GlueJob#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// The name you assign to the job definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#name GlueJob#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Non-overridable arguments for this job, specified as name-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#non_overridable_arguments GlueJob#non_overridable_arguments}
	NonOverridableArguments *string `field:"optional" json:"nonOverridableArguments" yaml:"nonOverridableArguments"`
	// Specifies configuration properties of a notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#notification_property GlueJob#notification_property}
	NotificationProperty *GlueJobNotificationProperty `field:"optional" json:"notificationProperty" yaml:"notificationProperty"`
	// The number of workers of a defined workerType that are allocated when a job runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#number_of_workers GlueJob#number_of_workers}
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The name of the SecurityConfiguration structure to be used with this job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#security_configuration GlueJob#security_configuration}
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The tags to use with this job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#tags GlueJob#tags}
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
	// The maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#timeout GlueJob#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// TThe type of predefined worker that is allocated when a job runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#worker_type GlueJob#worker_type}
	WorkerType *string `field:"optional" json:"workerType" yaml:"workerType"`
}

