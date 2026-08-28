// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens


type S3StorageLensStorageLensConfigurationDataExportS3BucketDestination struct {
	// The AWS account ID that owns the destination S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#account_id S3StorageLens#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The ARN of the bucket to which Amazon S3 Storage Lens exports will be placed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#arn S3StorageLens#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Configures the server-side encryption for Amazon S3 Storage Lens report files with either S3-managed keys (SSE-S3) or KMS-managed keys (SSE-KMS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#encryption S3StorageLens#encryption}
	Encryption *S3StorageLensStorageLensConfigurationDataExportS3BucketDestinationEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Specifies the file format to use when exporting Amazon S3 Storage Lens metrics export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#format S3StorageLens#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The version of the output schema to use when exporting Amazon S3 Storage Lens metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#output_schema_version S3StorageLens#output_schema_version}
	OutputSchemaVersion *string `field:"optional" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
	// The prefix to use for Amazon S3 Storage Lens export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#prefix S3StorageLens#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

