// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketNotificationConfigurationQueueConfigurationsFilterS3KeyRules struct {
	// The object key name prefix or suffix identifying one or more objects to which the filtering rule applies.
	//
	// The maximum length is 1,024 characters. Overlapping prefixes and suffixes are not supported. For more information, see [Configuring Event Notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3_bucket#name S3Bucket#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value that the filter searches for in object key names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3_bucket#value S3Bucket#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

