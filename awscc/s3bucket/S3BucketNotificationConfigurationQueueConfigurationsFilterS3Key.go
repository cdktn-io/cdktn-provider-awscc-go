// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketNotificationConfigurationQueueConfigurationsFilterS3Key struct {
	// A list of containers for the key-value pair that defines the criteria for the filter rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3_bucket#rules S3Bucket#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

