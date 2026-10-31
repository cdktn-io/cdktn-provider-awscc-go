// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotOtaUpdateConfig struct {
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
	// The files to be streamed by the OTA update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#files IotOtaUpdate#files}
	Files interface{} `field:"required" json:"files" yaml:"files"`
	// The ID of the OTA update to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#ota_update_id IotOtaUpdate#ota_update_id}
	OtaUpdateId *string `field:"required" json:"otaUpdateId" yaml:"otaUpdateId"`
	// The devices targeted to receive OTA updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#targets IotOtaUpdate#targets}
	Targets *[]*string `field:"required" json:"targets" yaml:"targets"`
	// A list of additional OTA update parameters, which are name-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#additional_parameters IotOtaUpdate#additional_parameters}
	AdditionalParameters *map[string]*string `field:"optional" json:"additionalParameters" yaml:"additionalParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#aws_job_abort_config IotOtaUpdate#aws_job_abort_config}.
	AwsJobAbortConfig *IotOtaUpdateAwsJobAbortConfig `field:"optional" json:"awsJobAbortConfig" yaml:"awsJobAbortConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#aws_job_executions_rollout_config IotOtaUpdate#aws_job_executions_rollout_config}.
	AwsJobExecutionsRolloutConfig *IotOtaUpdateAwsJobExecutionsRolloutConfig `field:"optional" json:"awsJobExecutionsRolloutConfig" yaml:"awsJobExecutionsRolloutConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#aws_job_presigned_url_config IotOtaUpdate#aws_job_presigned_url_config}.
	AwsJobPresignedUrlConfig *IotOtaUpdateAwsJobPresignedUrlConfig `field:"optional" json:"awsJobPresignedUrlConfig" yaml:"awsJobPresignedUrlConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#aws_job_timeout_config IotOtaUpdate#aws_job_timeout_config}.
	AwsJobTimeoutConfig *IotOtaUpdateAwsJobTimeoutConfig `field:"optional" json:"awsJobTimeoutConfig" yaml:"awsJobTimeoutConfig"`
	// The description of the OTA update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#description IotOtaUpdate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The protocol used to transfer the OTA update image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#protocols IotOtaUpdate#protocols}
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// The IAM role that grants AWS IoT Core access to the Amazon S3, IoT jobs and AWS Code Signing resources to create an OTA update job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#role_arn IotOtaUpdate#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Metadata which can be used to manage updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#tags IotOtaUpdate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether the update will continue to run (CONTINUOUS), or will be complete after all things specified as targets have completed the update (SNAPSHOT).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/iot_ota_update#target_selection IotOtaUpdate#target_selection}
	TargetSelection *string `field:"optional" json:"targetSelection" yaml:"targetSelection"`
}

