// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3outpostsbucket


type S3OutpostsBucketLifecycleConfigurationRules struct {
	// Specifies a lifecycle rule that stops incomplete multipart uploads to an Amazon S3Outposts bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3outposts_bucket#abort_incomplete_multipart_upload S3OutpostsBucket#abort_incomplete_multipart_upload}
	AbortIncompleteMultipartUpload *S3OutpostsBucketLifecycleConfigurationRulesAbortIncompleteMultipartUpload `field:"optional" json:"abortIncompleteMultipartUpload" yaml:"abortIncompleteMultipartUpload"`
	// Indicates when objects are deleted from Amazon S3Outposts.
	//
	// The date value must be in ISO 8601 format. The time is always midnight UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3outposts_bucket#expiration_date S3OutpostsBucket#expiration_date}
	ExpirationDate *string `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// Indicates the number of days after creation when objects are deleted from Amazon S3Outposts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3outposts_bucket#expiration_in_days S3OutpostsBucket#expiration_in_days}
	ExpirationInDays *float64 `field:"optional" json:"expirationInDays" yaml:"expirationInDays"`
	// The container for the filter of the lifecycle rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3outposts_bucket#filter S3OutpostsBucket#filter}
	Filter *S3OutpostsBucketLifecycleConfigurationRulesFilter `field:"optional" json:"filter" yaml:"filter"`
	// Unique identifier for the lifecycle rule. The value can't be longer than 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3outposts_bucket#id S3OutpostsBucket#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3outposts_bucket#status S3OutpostsBucket#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

