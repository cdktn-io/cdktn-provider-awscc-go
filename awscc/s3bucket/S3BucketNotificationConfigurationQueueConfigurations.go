// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketNotificationConfigurationQueueConfigurations struct {
	// The Amazon S3 bucket event about which you want to publish messages to Amazon SQS.
	//
	// For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#event S3Bucket#event}
	Event *string `field:"optional" json:"event" yaml:"event"`
	// The filtering rules that determine which objects trigger notifications.
	//
	// For example, you can create a filter so that Amazon S3 sends notifications only when image files with a ``.jpg`` extension are added to the bucket. For more information, see [Configuring event notifications using object key name filtering](https://docs.aws.amazon.com/AmazonS3/latest/user-guide/notification-how-to-filtering.html) in the *Amazon S3 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#filter S3Bucket#filter}
	Filter *S3BucketNotificationConfigurationQueueConfigurationsFilter `field:"optional" json:"filter" yaml:"filter"`
	// The Amazon Resource Name (ARN) of the Amazon SQS queue to which Amazon S3 publishes a message when it detects events of the specified type.
	//
	// FIFO queues are not allowed when enabling an SQS queue as the event notification destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#queue S3Bucket#queue}
	Queue *string `field:"optional" json:"queue" yaml:"queue"`
}

