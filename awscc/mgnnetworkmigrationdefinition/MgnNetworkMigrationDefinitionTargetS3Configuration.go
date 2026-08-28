// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mgnnetworkmigrationdefinition


type MgnNetworkMigrationDefinitionTargetS3Configuration struct {
	// The name of the S3 bucket for target artifacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mgn_network_migration_definition#s3_bucket MgnNetworkMigrationDefinition#s3_bucket}
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The AWS account ID of the S3 bucket owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mgn_network_migration_definition#s3_bucket_owner MgnNetworkMigrationDefinition#s3_bucket_owner}
	S3BucketOwner *string `field:"required" json:"s3BucketOwner" yaml:"s3BucketOwner"`
}

