// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelS3DestinationConfigurationDeadLetterQueueS3 struct {
	// The ARN of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/msk_channel#bucket_arn MskChannel#bucket_arn}
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// The error output prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/msk_channel#error_output_prefix MskChannel#error_output_prefix}
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// Optional 12-digit AWS account ID expected to own the dead-letter S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/msk_channel#expected_bucket_owner MskChannel#expected_bucket_owner}
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
}

