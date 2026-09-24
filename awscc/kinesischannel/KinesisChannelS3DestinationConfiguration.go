// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel


type KinesisChannelS3DestinationConfiguration struct {
	// The maximum time in seconds the channel buffers records before delivery if the minimum target file size is not reached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#data_freshness_in_seconds KinesisChannel#data_freshness_in_seconds}
	DataFreshnessInSeconds *float64 `field:"optional" json:"dataFreshnessInSeconds" yaml:"dataFreshnessInSeconds"`
	// Optional dead-letter queue (DLQ) configuration for records that cannot be delivered to the destination.
	//
	// When omitted, the service auto-fills using the storage BucketARN with an error prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#dead_letter_queue_s3_configuration KinesisChannel#dead_letter_queue_s3_configuration}
	DeadLetterQueueS3Configuration *KinesisChannelS3DestinationConfigurationDeadLetterQueueS3Configuration `field:"optional" json:"deadLetterQueueS3Configuration" yaml:"deadLetterQueueS3Configuration"`
	// S3 storage configuration including the destination bucket, output key template, storage class, and compression type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#storage_configuration KinesisChannel#storage_configuration}
	StorageConfiguration *KinesisChannelS3DestinationConfigurationStorageConfiguration `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
}

