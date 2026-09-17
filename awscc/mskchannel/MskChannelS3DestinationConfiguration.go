// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelS3DestinationConfiguration struct {
	// Data freshness in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#data_freshness_in_seconds MskChannel#data_freshness_in_seconds}
	DataFreshnessInSeconds *float64 `field:"optional" json:"dataFreshnessInSeconds" yaml:"dataFreshnessInSeconds"`
	// Dead letter queue S3 configuration of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#dead_letter_queue_s3 MskChannel#dead_letter_queue_s3}
	DeadLetterQueueS3 *MskChannelS3DestinationConfigurationDeadLetterQueueS3 `field:"optional" json:"deadLetterQueueS3" yaml:"deadLetterQueueS3"`
	// The Amazon Resource Name (ARN) of an IAM role used by MSK to access S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#service_execution_role_arn MskChannel#service_execution_role_arn}
	ServiceExecutionRoleArn *string `field:"optional" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// S3 storage configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#storage MskChannel#storage}
	Storage *MskChannelS3DestinationConfigurationStorage `field:"optional" json:"storage" yaml:"storage"`
}

