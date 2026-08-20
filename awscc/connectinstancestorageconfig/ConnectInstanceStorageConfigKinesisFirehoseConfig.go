// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectinstancestorageconfig


type ConnectInstanceStorageConfigKinesisFirehoseConfig struct {
	// An ARN is a unique AWS resource identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_instance_storage_config#firehose_arn ConnectInstanceStorageConfig#firehose_arn}
	FirehoseArn *string `field:"optional" json:"firehoseArn" yaml:"firehoseArn"`
}

