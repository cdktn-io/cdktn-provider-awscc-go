// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjobtemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotJobTemplateConfig struct {
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
	// A description of the Job Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job_template#description IotJobTemplate#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job_template#job_template_id IotJobTemplate#job_template_id}.
	JobTemplateId *string `field:"required" json:"jobTemplateId" yaml:"jobTemplateId"`
	// The criteria that determine when and how a job abort takes place.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job_template#abort_config IotJobTemplate#abort_config}
	AbortConfig *IotJobTemplateAbortConfig `field:"optional" json:"abortConfig" yaml:"abortConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job_template#destination_package_versions IotJobTemplate#destination_package_versions}.
	DestinationPackageVersions *[]*string `field:"optional" json:"destinationPackageVersions" yaml:"destinationPackageVersions"`
	// The job document. Required if you don't specify a value for documentSource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job_template#document IotJobTemplate#document}
	Document *string `field:"optional" json:"document" yaml:"document"`
	// An S3 link to the job document to use in the template.
	//
	// Required if you don't specify a value for document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job_template#document_source IotJobTemplate#document_source}
	DocumentSource *string `field:"optional" json:"documentSource" yaml:"documentSource"`
	// Optional for copying a JobTemplate from a pre-existing Job configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job_template#job_arn IotJobTemplate#job_arn}
	JobArn *string `field:"optional" json:"jobArn" yaml:"jobArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job_template#job_executions_retry_config IotJobTemplate#job_executions_retry_config}.
	JobExecutionsRetryConfig *IotJobTemplateJobExecutionsRetryConfig `field:"optional" json:"jobExecutionsRetryConfig" yaml:"jobExecutionsRetryConfig"`
	// Allows you to create a staged rollout of a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job_template#job_executions_rollout_config IotJobTemplate#job_executions_rollout_config}
	JobExecutionsRolloutConfig *IotJobTemplateJobExecutionsRolloutConfig `field:"optional" json:"jobExecutionsRolloutConfig" yaml:"jobExecutionsRolloutConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job_template#maintenance_windows IotJobTemplate#maintenance_windows}.
	MaintenanceWindows interface{} `field:"optional" json:"maintenanceWindows" yaml:"maintenanceWindows"`
	// Configuration for pre-signed S3 URLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job_template#presigned_url_config IotJobTemplate#presigned_url_config}
	PresignedUrlConfig *IotJobTemplatePresignedUrlConfig `field:"optional" json:"presignedUrlConfig" yaml:"presignedUrlConfig"`
	// Metadata that can be used to manage the JobTemplate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job_template#tags IotJobTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the amount of time each device has to finish its execution of the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job_template#timeout_config IotJobTemplate#timeout_config}
	TimeoutConfig *IotJobTemplateTimeoutConfig `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
}

