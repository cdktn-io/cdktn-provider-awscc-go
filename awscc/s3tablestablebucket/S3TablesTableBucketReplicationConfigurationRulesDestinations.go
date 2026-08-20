// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestablebucket


type S3TablesTableBucketReplicationConfigurationRulesDestinations struct {
	// The ARN of the destination table bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3tables_table_bucket#destination_table_bucket_arn S3TablesTableBucket#destination_table_bucket_arn}
	DestinationTableBucketArn *string `field:"optional" json:"destinationTableBucketArn" yaml:"destinationTableBucketArn"`
}

