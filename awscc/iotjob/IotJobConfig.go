// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjob

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotJobConfig struct {
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
	// A job identifier which must be unique for your AWS account.
	//
	// We recommend using a UUID. Alpha-numeric characters, '-' and '_' are valid for use here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#job_id IotJob#job_id}
	JobId *string `field:"required" json:"jobId" yaml:"jobId"`
	// A list of things and thing groups to which the job should be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#targets IotJob#targets}
	Targets *[]*string `field:"required" json:"targets" yaml:"targets"`
	// The criteria that determine when and how a job abort takes place.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#abort_config IotJob#abort_config}
	AbortConfig *IotJobAbortConfig `field:"optional" json:"abortConfig" yaml:"abortConfig"`
	// A short text description of the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#description IotJob#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The package version Amazon Resource Names (ARNs) that are installed on the device when the job successfully completes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#destination_package_versions IotJob#destination_package_versions}
	DestinationPackageVersions *[]*string `field:"optional" json:"destinationPackageVersions" yaml:"destinationPackageVersions"`
	// The job document. Required if you don't specify a value for documentSource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#document IotJob#document}
	Document *string `field:"optional" json:"document" yaml:"document"`
	// Parameters of an Amazon Web Services managed template that you can specify to create the job document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#document_parameters IotJob#document_parameters}
	DocumentParameters *map[string]*string `field:"optional" json:"documentParameters" yaml:"documentParameters"`
	// An S3 link, or S3 object URL, to the job document.
	//
	// The link is an Amazon S3 object URL and is required if you don't specify a value for document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#document_source IotJob#document_source}
	DocumentSource *string `field:"optional" json:"documentSource" yaml:"documentSource"`
	// The configuration that determines how many retries are allowed for each failure type for a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#job_executions_retry_config IotJob#job_executions_retry_config}
	JobExecutionsRetryConfig *IotJobJobExecutionsRetryConfig `field:"optional" json:"jobExecutionsRetryConfig" yaml:"jobExecutionsRetryConfig"`
	// Allows you to create a staged rollout of a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#job_executions_rollout_config IotJob#job_executions_rollout_config}
	JobExecutionsRolloutConfig *IotJobJobExecutionsRolloutConfig `field:"optional" json:"jobExecutionsRolloutConfig" yaml:"jobExecutionsRolloutConfig"`
	// The ARN of the job template used to create the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#job_template_arn IotJob#job_template_arn}
	JobTemplateArn *string `field:"optional" json:"jobTemplateArn" yaml:"jobTemplateArn"`
	// Configuration for pre-signed S3 URLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#presigned_url_config IotJob#presigned_url_config}
	PresignedUrlConfig *IotJobPresignedUrlConfig `field:"optional" json:"presignedUrlConfig" yaml:"presignedUrlConfig"`
	// Specifies the date and time that a job will begin the rollout of the job document to all devices in the target group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#scheduling_config IotJob#scheduling_config}
	SchedulingConfig *IotJobSchedulingConfig `field:"optional" json:"schedulingConfig" yaml:"schedulingConfig"`
	// Metadata which can be used to manage the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#tags IotJob#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether the job will continue to run (CONTINUOUS), or will be complete after all those things specified as targets have completed the job (SNAPSHOT).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#target_selection IotJob#target_selection}
	TargetSelection *string `field:"optional" json:"targetSelection" yaml:"targetSelection"`
	// Specifies the amount of time each device has to finish its execution of the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_job#timeout_config IotJob#timeout_config}
	TimeoutConfig *IotJobTimeoutConfig `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
}

