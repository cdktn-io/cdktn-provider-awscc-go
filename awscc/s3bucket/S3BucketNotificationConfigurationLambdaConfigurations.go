// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketNotificationConfigurationLambdaConfigurations struct {
	// The Amazon S3 bucket event for which to invoke the LAMlong function.
	//
	// For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#event S3Bucket#event}
	Event *string `field:"optional" json:"event" yaml:"event"`
	// The filtering rules that determine which objects invoke the AWS Lambda function.
	//
	// For example, you can create a filter so that only image files with a ``.jpg`` extension invoke the function when they are added to the Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#filter S3Bucket#filter}
	Filter *S3BucketNotificationConfigurationLambdaConfigurationsFilter `field:"optional" json:"filter" yaml:"filter"`
	// The Amazon Resource Name (ARN) of the LAMlong function that Amazon S3 invokes when the specified event type occurs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#function S3Bucket#function}
	Function *string `field:"optional" json:"function" yaml:"function"`
}

