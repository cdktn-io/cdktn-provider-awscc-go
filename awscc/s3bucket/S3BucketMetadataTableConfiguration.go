// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketMetadataTableConfiguration struct {
	// The destination information for the metadata table configuration.
	//
	// The destination table bucket must be in the same Region and AWS-account as the general purpose bucket. The specified metadata table name must be unique within the ``aws_s3_metadata`` namespace in the destination table bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#s3_tables_destination S3Bucket#s3_tables_destination}
	S3TablesDestination *S3BucketMetadataTableConfigurationS3TablesDestination `field:"optional" json:"s3TablesDestination" yaml:"s3TablesDestination"`
}

