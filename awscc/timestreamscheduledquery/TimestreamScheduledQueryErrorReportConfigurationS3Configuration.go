// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamscheduledquery


type TimestreamScheduledQueryErrorReportConfigurationS3Configuration struct {
	// Name of the S3 bucket under which error reports will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/timestream_scheduled_query#bucket_name TimestreamScheduledQuery#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Encryption at rest options for the error reports.
	//
	// If no encryption option is specified, Timestream will choose SSE_S3 as default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/timestream_scheduled_query#encryption_option TimestreamScheduledQuery#encryption_option}
	EncryptionOption *string `field:"optional" json:"encryptionOption" yaml:"encryptionOption"`
	// Prefix for error report keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/timestream_scheduled_query#object_key_prefix TimestreamScheduledQuery#object_key_prefix}
	ObjectKeyPrefix *string `field:"optional" json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
}

