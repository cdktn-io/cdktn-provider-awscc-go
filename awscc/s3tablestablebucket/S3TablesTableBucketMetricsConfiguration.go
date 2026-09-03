// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestablebucket


type S3TablesTableBucketMetricsConfiguration struct {
	// Indicates whether Metrics are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3tables_table_bucket#status S3TablesTableBucket#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

