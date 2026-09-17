// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketLifecycleConfigurationRulesAbortIncompleteMultipartUpload struct {
	// Specifies the number of days after which Amazon S3 stops an incomplete multipart upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#days_after_initiation S3Bucket#days_after_initiation}
	DaysAfterInitiation *float64 `field:"optional" json:"daysAfterInitiation" yaml:"daysAfterInitiation"`
}

