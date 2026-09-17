// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel


type KinesisChannelStreamConfigurationListRecordConfiguration struct {
	// The format used to interpret records read from the source stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesis_channel#record_format_type KinesisChannel#record_format_type}
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
	// The ARN of the AWS Glue Schema Registry (GSR) schema.
	//
	// Required for the S3 Tables destination, where it is used to create the S3 Table and to validate that the record format matches the table schema. Also used when RecordFormatType is GSR_JSON to interpret records read from the source stream. Vanilla S3 delivery writes records as S3 objects and does not need a schema. The schema must be in the same account and region as the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesis_channel#gsr_schema_arn KinesisChannel#gsr_schema_arn}
	GsrSchemaArn *string `field:"optional" json:"gsrSchemaArn" yaml:"gsrSchemaArn"`
}

