// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketNotificationConfigurationQueueConfigurationsFilter struct {
	// A container for object key name prefix and suffix filtering rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_bucket#s3_key S3Bucket#s3_key}
	S3Key *S3BucketNotificationConfigurationQueueConfigurationsFilterS3Key `field:"optional" json:"s3Key" yaml:"s3Key"`
}

