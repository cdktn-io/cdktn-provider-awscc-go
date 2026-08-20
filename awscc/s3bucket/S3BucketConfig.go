// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3BucketConfig struct {
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
	// Configures the transfer acceleration state for an Amazon S3 bucket.
	//
	// For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#accelerate_configuration S3Bucket#accelerate_configuration}
	AccelerateConfiguration *S3BucketAccelerateConfiguration `field:"optional" json:"accelerateConfiguration" yaml:"accelerateConfiguration"`
	// This is a legacy property, and it is not recommended for most use cases.
	//
	// A majority of modern use cases in Amazon S3 no longer require the use of ACLs, and we recommend that you keep ACLs disabled. For more information, see [Controlling object ownership](https://docs.aws.amazon.com//AmazonS3/latest/userguide/about-object-ownership.html) in the *Amazon S3 User Guide*.
	//   A canned access control list (ACL) that grants predefined permissions to the bucket. For more information about canned ACLs, see [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) in the *Amazon S3 User Guide*.
	//   S3 buckets are created with ACLs disabled by default. Therefore, unless you explicitly set the [AWS::S3::OwnershipControls](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ownershipcontrols.html) property to enable ACLs, your resource will fail to deploy with any value other than Private. Use cases requiring ACLs are uncommon.
	//   The majority of access control configurations can be successfully and more easily achieved with bucket policies. For more information, see [AWS::S3::BucketPolicy](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html). For examples of common policy configurations, including S3 Server Access Logs buckets and more, see [Bucket policy examples](https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html) in the *Amazon S3 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#access_control S3Bucket#access_control}
	AccessControl *string `field:"optional" json:"accessControl" yaml:"accessControl"`
	// Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#analytics_configurations S3Bucket#analytics_configurations}
	AnalyticsConfigurations interface{} `field:"optional" json:"analyticsConfigurations" yaml:"analyticsConfigurations"`
	// Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3), AWS KMS-managed keys (SSE-KMS), or dual-layer server-side encryption with KMS-managed keys (DSSE-KMS).
	//
	// For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#bucket_encryption S3Bucket#bucket_encryption}
	BucketEncryption *S3BucketBucketEncryption `field:"optional" json:"bucketEncryption" yaml:"bucketEncryption"`
	// A name for the bucket.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html). For more information, see [Rules for naming Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html) in the *Amazon S3 User Guide*.
	//   If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#bucket_name S3Bucket#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Describes the cross-origin access configuration for objects in an Amazon S3 bucket.
	//
	// For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#cors_configuration S3Bucket#cors_configuration}
	CorsConfiguration *S3BucketCorsConfiguration `field:"optional" json:"corsConfiguration" yaml:"corsConfiguration"`
	// Defines how Amazon S3 handles Intelligent-Tiering storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#intelligent_tiering_configurations S3Bucket#intelligent_tiering_configurations}
	IntelligentTieringConfigurations interface{} `field:"optional" json:"intelligentTieringConfigurations" yaml:"intelligentTieringConfigurations"`
	// Specifies the inventory configuration for an Amazon S3 bucket.
	//
	// For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#inventory_configurations S3Bucket#inventory_configurations}
	InventoryConfigurations interface{} `field:"optional" json:"inventoryConfigurations" yaml:"inventoryConfigurations"`
	// Specifies the lifecycle configuration for objects in an Amazon S3 bucket.
	//
	// For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#lifecycle_configuration S3Bucket#lifecycle_configuration}
	LifecycleConfiguration *S3BucketLifecycleConfiguration `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// Settings that define where logs are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#logging_configuration S3Bucket#logging_configuration}
	LoggingConfiguration *S3BucketLoggingConfiguration `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// The metadata table configuration of an S3 general purpose bucket.
	//
	// For more information, see [Accelerating data discovery with S3 Metadata](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html) and [Setting up permissions for configuring metadata tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#metadata_table_configuration S3Bucket#metadata_table_configuration}
	MetadataTableConfiguration *S3BucketMetadataTableConfiguration `field:"optional" json:"metadataTableConfiguration" yaml:"metadataTableConfiguration"`
	// Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket.
	//
	// If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#metrics_configurations S3Bucket#metrics_configurations}
	MetricsConfigurations interface{} `field:"optional" json:"metricsConfigurations" yaml:"metricsConfigurations"`
	// Configuration that defines how Amazon S3 handles bucket notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#notification_configuration S3Bucket#notification_configuration}
	NotificationConfiguration *S3BucketNotificationConfiguration `field:"optional" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// This operation is not supported for directory buckets.
	//
	// Places an Object Lock configuration on the specified bucket. The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
	//    +  The ``DefaultRetention`` settings require both a mode and a period.
	//   +  The ``DefaultRetention`` period can be either ``Days`` or ``Years`` but you must select one. You cannot specify ``Days`` and ``Years`` at the same time.
	//   +  You can enable Object Lock for new or existing buckets. For more information, see [Configuring Object Lock](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-configure.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#object_lock_configuration S3Bucket#object_lock_configuration}
	ObjectLockConfiguration *S3BucketObjectLockConfiguration `field:"optional" json:"objectLockConfiguration" yaml:"objectLockConfiguration"`
	// Indicates whether this bucket has an Object Lock configuration enabled. Enable ``ObjectLockEnabled`` when you apply ``ObjectLockConfiguration`` to a bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#object_lock_enabled S3Bucket#object_lock_enabled}
	ObjectLockEnabled interface{} `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// Configuration that defines how Amazon S3 handles Object Ownership rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#ownership_controls S3Bucket#ownership_controls}
	OwnershipControls *S3BucketOwnershipControls `field:"optional" json:"ownershipControls" yaml:"ownershipControls"`
	// Configuration that defines how Amazon S3 handles public access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#public_access_block_configuration S3Bucket#public_access_block_configuration}
	PublicAccessBlockConfiguration *S3BucketPublicAccessBlockConfiguration `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
	// Configuration for replicating objects in an S3 bucket.
	//
	// To enable replication, you must also enable versioning by using the ``VersioningConfiguration`` property.
	//  Amazon S3 can store replicated objects in a single destination bucket or multiple destination buckets. The destination bucket or buckets must already exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#replication_configuration S3Bucket#replication_configuration}
	ReplicationConfiguration *S3BucketReplicationConfiguration `field:"optional" json:"replicationConfiguration" yaml:"replicationConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#tags S3Bucket#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Enables multiple versions of all objects in this bucket.
	//
	// You might enable versioning to prevent objects from being deleted or overwritten by mistake or to archive objects so that you can retrieve previous versions of them.
	//   When you enable versioning on a bucket for the first time, it might take a short amount of time for the change to be fully propagated. We recommend that you wait for 15 minutes after enabling versioning before issuing write operations (``PUT`` or ``DELETE``) on objects in the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#versioning_configuration S3Bucket#versioning_configuration}
	VersioningConfiguration *S3BucketVersioningConfiguration `field:"optional" json:"versioningConfiguration" yaml:"versioningConfiguration"`
	// Information used to configure the bucket as a static website. For more information, see [Hosting Websites on Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#website_configuration S3Bucket#website_configuration}
	WebsiteConfiguration *S3BucketWebsiteConfiguration `field:"optional" json:"websiteConfiguration" yaml:"websiteConfiguration"`
}

