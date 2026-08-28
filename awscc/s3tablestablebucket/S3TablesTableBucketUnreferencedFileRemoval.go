// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestablebucket


type S3TablesTableBucketUnreferencedFileRemoval struct {
	// S3 permanently deletes noncurrent objects after the number of days specified by the NoncurrentDays property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3tables_table_bucket#noncurrent_days S3TablesTableBucket#noncurrent_days}
	NoncurrentDays *float64 `field:"optional" json:"noncurrentDays" yaml:"noncurrentDays"`
	// Indicates whether the Unreferenced File Removal maintenance action is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3tables_table_bucket#status S3TablesTableBucket#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// For any object not referenced by your table and older than the UnreferencedDays property, S3 creates a delete marker and marks the object version as noncurrent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3tables_table_bucket#unreferenced_days S3TablesTableBucket#unreferenced_days}
	UnreferencedDays *float64 `field:"optional" json:"unreferencedDays" yaml:"unreferencedDays"`
}

