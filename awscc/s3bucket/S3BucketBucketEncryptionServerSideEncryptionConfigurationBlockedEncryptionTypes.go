// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketBucketEncryptionServerSideEncryptionConfigurationBlockedEncryptionTypes struct {
	// The object encryption type that you want to block or unblock for an Amazon S3 general purpose bucket.
	//
	// Currently, this parameter only supports blocking or unblocking server side encryption with customer-provided keys (SSE-C). For more information about SSE-C, see [Using server-side encryption with customer-provided keys (SSE-C)](https://docs.aws.amazon.com/AmazonS3/latest/userguide/ServerSideEncryptionCustomerKeys.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/s3_bucket#encryption_type S3Bucket#encryption_type}
	EncryptionType *[]*string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
}

