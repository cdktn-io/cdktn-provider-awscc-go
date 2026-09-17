// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestablebucket


type S3TablesTableBucketReplicationConfiguration struct {
	// The ARN of the IAM role to use for replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3tables_table_bucket#role S3TablesTableBucket#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// List of replication rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3tables_table_bucket#rules S3TablesTableBucket#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

