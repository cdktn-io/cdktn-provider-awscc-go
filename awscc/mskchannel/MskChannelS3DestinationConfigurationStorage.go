// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelS3DestinationConfigurationStorage struct {
	// ARN of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_channel#bucket_arn MskChannel#bucket_arn}
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// S3 compression type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_channel#compression_type MskChannel#compression_type}
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Optional 12-digit AWS account ID expected to own the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_channel#expected_bucket_owner MskChannel#expected_bucket_owner}
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Template for S3 key for output objects, used for partitioning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_channel#output_key_template MskChannel#output_key_template}
	OutputKeyTemplate *string `field:"optional" json:"outputKeyTemplate" yaml:"outputKeyTemplate"`
	// Optional prefix for output objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_channel#output_prefix MskChannel#output_prefix}
	OutputPrefix *string `field:"optional" json:"outputPrefix" yaml:"outputPrefix"`
	// S3 storage class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_channel#storage_class MskChannel#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

