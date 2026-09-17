// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfigurationRulesDestination struct {
	// Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS-account that owns the destination bucket.
	//
	// If this is not specified in the replication configuration, the replicas are owned by same AWS-account that owns the source object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#access_control_translation S3Bucket#access_control_translation}
	AccessControlTranslation *S3BucketReplicationConfigurationRulesDestinationAccessControlTranslation `field:"optional" json:"accessControlTranslation" yaml:"accessControlTranslation"`
	// Destination bucket owner account ID.
	//
	// In a cross-account scenario, if you direct Amazon S3 to change replica ownership to the AWS-account that owns the destination bucket by specifying the ``AccessControlTranslation`` property, this is the account ID of the destination bucket owner. For more information, see [Cross-Region Replication Additional Configuration: Change Replica Owner](https://docs.aws.amazon.com/AmazonS3/latest/dev/crr-change-owner.html) in the *Amazon S3 User Guide*.
	//  If you specify the ``AccessControlTranslation`` property, the ``Account`` property is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#account S3Bucket#account}
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The Amazon Resource Name (ARN) of the bucket where you want Amazon S3 to store the results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#bucket S3Bucket#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Specifies encryption-related information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#encryption_configuration S3Bucket#encryption_configuration}
	EncryptionConfiguration *S3BucketReplicationConfigurationRulesDestinationEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// A container specifying replication metrics-related settings enabling replication metrics and events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#metrics S3Bucket#metrics}
	Metrics *S3BucketReplicationConfigurationRulesDestinationMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	// A container specifying S3 Replication Time Control (S3 RTC), including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated.
	//
	// Must be specified together with a ``Metrics`` block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#replication_time S3Bucket#replication_time}
	ReplicationTime *S3BucketReplicationConfigurationRulesDestinationReplicationTime `field:"optional" json:"replicationTime" yaml:"replicationTime"`
	// The storage class to use when replicating objects, such as S3 Standard or reduced redundancy.
	//
	// By default, Amazon S3 uses the storage class of the source object to create the object replica.
	//  For valid values, see the ``StorageClass`` element of the [PUT Bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) action in the *Amazon S3 API Reference*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#storage_class S3Bucket#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

