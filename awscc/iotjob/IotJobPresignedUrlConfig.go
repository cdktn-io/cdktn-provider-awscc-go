// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjob


type IotJobPresignedUrlConfig struct {
	// How long (in seconds) pre-signed URLs are valid.
	//
	// Valid values are 60 - 3600, the default value is 3600 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job#expires_in_sec IotJob#expires_in_sec}
	ExpiresInSec *float64 `field:"optional" json:"expiresInSec" yaml:"expiresInSec"`
	// The ARN of an IAM role that grants permission to download files from the S3 bucket where the job data/updates are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job#role_arn IotJob#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

