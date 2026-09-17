// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel


type KinesisChannelS3DestinationConfigurationStorageConfiguration struct {
	// The ARN of the S3 bucket for record delivery.
	//
	// Different channels can deliver to the same bucket. Buckets can be cross-account but must be in the same region as the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesis_channel#bucket_arn KinesisChannel#bucket_arn}
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// The compression algorithm applied to delivered objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesis_channel#compression_type KinesisChannel#compression_type}
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// The AWS account ID of the expected owner of the destination S3 bucket.
	//
	// Used to verify bucket ownership before delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesis_channel#expected_bucket_owner KinesisChannel#expected_bucket_owner}
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Optional template for the S3 object key path.
	//
	// Supports placeholders in the form !{name}: !{channel-name}, !{channel-id}, !{stream-name}, !{yyyy}, !{yy}, !{MM}, !{dd}, !{HH}, !{mm}, and !{extension} (a literal file extension can be supplied as !{extension:.json.gz}). When omitted, the service uses the default 'kinesis-channel/!{channel-name}/!{channel-id}/!{yyyy}/!{MM}/!{dd}/!{HH}/!{channel-name}-!{channel-id}-!{yyyy}-!{MM}-!{dd}-!{HH}-!{mm}!{extension}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesis_channel#output_key_template KinesisChannel#output_key_template}
	OutputKeyTemplate *string `field:"optional" json:"outputKeyTemplate" yaml:"outputKeyTemplate"`
	// The S3 storage class for delivered objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesis_channel#storage_class KinesisChannel#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

