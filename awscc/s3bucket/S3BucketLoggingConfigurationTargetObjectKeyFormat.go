// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketLoggingConfigurationTargetObjectKeyFormat struct {
	// Amazon S3 keys for log objects are partitioned in the following format:   ``[DestinationPrefix][SourceAccountId]/[SourceRegion]/[SourceBucket]/[YYYY]/[MM]/[DD]/[YYYY]-[MM]-[DD]-[hh]-[mm]-[ss]-[UniqueString]``   PartitionedPrefix defaults to EventTime delivery when server access logs are delivered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3_bucket#partitioned_prefix S3Bucket#partitioned_prefix}
	PartitionedPrefix *S3BucketLoggingConfigurationTargetObjectKeyFormatPartitionedPrefix `field:"optional" json:"partitionedPrefix" yaml:"partitionedPrefix"`
	// This format defaults the prefix to the given log file prefix for delivering server access log file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3_bucket#simple_prefix S3Bucket#simple_prefix}
	SimplePrefix *string `field:"optional" json:"simplePrefix" yaml:"simplePrefix"`
}

