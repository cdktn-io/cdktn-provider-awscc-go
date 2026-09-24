// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel


type KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStruct struct {
	// The compression algorithm applied to objects delivered to the S3 Tables destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#compression_type KinesisChannel#compression_type}
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// The name of the S3 Tables namespace that contains the destination table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#namespace KinesisChannel#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The partition specification used by the destination Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#partition_spec KinesisChannel#partition_spec}
	PartitionSpec *KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListPartitionSpec `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
	// The ARN of the S3 Tables table bucket for record delivery.
	//
	// Buckets can be cross-account but must be in the same region as the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#table_bucket_arn KinesisChannel#table_bucket_arn}
	TableBucketArn *string `field:"optional" json:"tableBucketArn" yaml:"tableBucketArn"`
	// The name of the destination S3 Tables table.
	//
	// The table is created for the customer if it does not yet exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#table_name KinesisChannel#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

