// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketLoggingConfiguration struct {
	// The name of the bucket where Amazon S3 should store server access log files.
	//
	// You can store log files in any bucket that you own. By default, logs are stored in the bucket where the ``LoggingConfiguration`` property is defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3_bucket#destination_bucket_name S3Bucket#destination_bucket_name}
	DestinationBucketName *string `field:"optional" json:"destinationBucketName" yaml:"destinationBucketName"`
	// A prefix for all log object keys.
	//
	// If you store log files from multiple Amazon S3 buckets in a single bucket, you can use a prefix to distinguish which log files came from which bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3_bucket#log_file_prefix S3Bucket#log_file_prefix}
	LogFilePrefix *string `field:"optional" json:"logFilePrefix" yaml:"logFilePrefix"`
	// Amazon S3 key format for log objects. Only one format, either PartitionedPrefix or SimplePrefix, is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3_bucket#target_object_key_format S3Bucket#target_object_key_format}
	TargetObjectKeyFormat *S3BucketLoggingConfigurationTargetObjectKeyFormat `field:"optional" json:"targetObjectKeyFormat" yaml:"targetObjectKeyFormat"`
}

