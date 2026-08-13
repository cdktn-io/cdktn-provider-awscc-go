// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjobtemplate


type IotJobTemplatePresignedUrlConfig struct {
	// How number (in seconds) pre-signed URLs are valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_job_template#expires_in_sec IotJobTemplate#expires_in_sec}
	ExpiresInSec *float64 `field:"optional" json:"expiresInSec" yaml:"expiresInSec"`
	// The ARN of an IAM role that grants grants permission to download files from the S3 bucket where the job data/updates are stored.
	//
	// The role must also grant permission for IoT to download the files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_job_template#role_arn IotJobTemplate#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

