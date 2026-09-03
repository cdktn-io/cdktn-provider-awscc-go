// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketMetadataTableConfigurationS3TablesDestination struct {
	// The Amazon Resource Name (ARN) for the table bucket that's specified as the destination in the metadata table configuration.
	//
	// The destination table bucket must be in the same Region and AWS-account as the general purpose bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3_bucket#table_bucket_arn S3Bucket#table_bucket_arn}
	TableBucketArn *string `field:"optional" json:"tableBucketArn" yaml:"tableBucketArn"`
	// The name for the metadata table in your metadata table configuration.
	//
	// The specified metadata table name must be unique within the ``aws_s3_metadata`` namespace in the destination table bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3_bucket#table_name S3Bucket#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

