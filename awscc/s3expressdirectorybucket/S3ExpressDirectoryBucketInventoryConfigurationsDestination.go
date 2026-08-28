// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3expressdirectorybucket


type S3ExpressDirectoryBucketInventoryConfigurationsDestination struct {
	// The account ID that owns the destination S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3express_directory_bucket#bucket_account_id S3ExpressDirectoryBucket#bucket_account_id}
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
	// The Amazon Resource Name (ARN) of the destination Amazon S3 bucket to which data is exported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3express_directory_bucket#bucket_arn S3ExpressDirectoryBucket#bucket_arn}
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// Specifies the file format used when exporting data to Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3express_directory_bucket#format S3ExpressDirectoryBucket#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The prefix to use when exporting data. The prefix is prepended to all results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3express_directory_bucket#prefix S3ExpressDirectoryBucket#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

