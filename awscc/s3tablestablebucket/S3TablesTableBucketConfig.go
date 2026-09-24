// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3tablestablebucket

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3TablesTableBucketConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A name for the table bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table_bucket#table_bucket_name S3TablesTableBucket#table_bucket_name}
	TableBucketName *string `field:"required" json:"tableBucketName" yaml:"tableBucketName"`
	// Specifies encryption settings for the table bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table_bucket#encryption_configuration S3TablesTableBucket#encryption_configuration}
	EncryptionConfiguration *S3TablesTableBucketEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Settings governing the Metric configuration for the table bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table_bucket#metrics_configuration S3TablesTableBucket#metrics_configuration}
	MetricsConfiguration *S3TablesTableBucketMetricsConfiguration `field:"optional" json:"metricsConfiguration" yaml:"metricsConfiguration"`
	// Specifies replication configuration for the table bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table_bucket#replication_configuration S3TablesTableBucket#replication_configuration}
	ReplicationConfiguration *S3TablesTableBucketReplicationConfiguration `field:"optional" json:"replicationConfiguration" yaml:"replicationConfiguration"`
	// Specifies storage class settings for the table bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table_bucket#storage_class_configuration S3TablesTableBucket#storage_class_configuration}
	StorageClassConfiguration *S3TablesTableBucketStorageClassConfiguration `field:"optional" json:"storageClassConfiguration" yaml:"storageClassConfiguration"`
	// User tags (key-value pairs) to associate with the table bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table_bucket#tags S3TablesTableBucket#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Settings governing the Unreferenced File Removal maintenance action.
	//
	// Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3tables_table_bucket#unreferenced_file_removal S3TablesTableBucket#unreferenced_file_removal}
	UnreferencedFileRemoval *S3TablesTableBucketUnreferencedFileRemoval `field:"optional" json:"unreferencedFileRemoval" yaml:"unreferencedFileRemoval"`
}

