// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestablebucket


type S3TablesTableBucketReplicationConfigurationRules struct {
	// List of replication destinations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3tables_table_bucket#destinations S3TablesTableBucket#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
}

