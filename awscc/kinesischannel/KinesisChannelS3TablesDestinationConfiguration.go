// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel


type KinesisChannelS3TablesDestinationConfiguration struct {
	// The maximum time in seconds the channel buffers records before delivery if the minimum target file size is not reached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#data_freshness_in_seconds KinesisChannel#data_freshness_in_seconds}
	DataFreshnessInSeconds *float64 `field:"optional" json:"dataFreshnessInSeconds" yaml:"dataFreshnessInSeconds"`
	// The dead-letter queue (DLQ) configuration for records that cannot be delivered to the S3 Tables destination.
	//
	// Required for S3 Tables: there is no safe fallback because S3 Tables metadata writes are critical-path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#dead_letter_queue_s3_configuration KinesisChannel#dead_letter_queue_s3_configuration}
	DeadLetterQueueS3Configuration *KinesisChannelS3TablesDestinationConfigurationDeadLetterQueueS3Configuration `field:"optional" json:"deadLetterQueueS3Configuration" yaml:"deadLetterQueueS3Configuration"`
	// The list of S3 Tables destinations.
	//
	// v1 supports a single element; the list shape allows future extensibility to fan out to multiple tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#s3_tables_configuration_list KinesisChannel#s3_tables_configuration_list}
	S3TablesConfigurationList interface{} `field:"optional" json:"s3TablesConfigurationList" yaml:"s3TablesConfigurationList"`
}

