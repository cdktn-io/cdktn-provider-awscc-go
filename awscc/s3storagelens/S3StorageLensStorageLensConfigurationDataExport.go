// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens


type S3StorageLensStorageLensConfigurationDataExport struct {
	// CloudWatch metrics settings for the Amazon S3 Storage Lens metrics export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_storage_lens#cloudwatch_metrics S3StorageLens#cloudwatch_metrics}
	CloudwatchMetrics *S3StorageLensStorageLensConfigurationDataExportCloudwatchMetrics `field:"optional" json:"cloudwatchMetrics" yaml:"cloudwatchMetrics"`
	// S3 bucket destination settings for the Amazon S3 Storage Lens metrics export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_storage_lens#s3_bucket_destination S3StorageLens#s3_bucket_destination}
	S3BucketDestination *S3StorageLensStorageLensConfigurationDataExportS3BucketDestination `field:"optional" json:"s3BucketDestination" yaml:"s3BucketDestination"`
	// S3 Tables destination settings for the Amazon S3 Storage Lens metrics export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_storage_lens#storage_lens_table_destination S3StorageLens#storage_lens_table_destination}
	StorageLensTableDestination *S3StorageLensStorageLensConfigurationDataExportStorageLensTableDestination `field:"optional" json:"storageLensTableDestination" yaml:"storageLensTableDestination"`
}

