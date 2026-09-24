// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketLoggingConfigurationTargetObjectKeyFormatPartitionedPrefix struct {
	// Specifies the partition date source for the partitioned prefix.
	//
	// ``PartitionDateSource`` can be ``EventTime`` or ``DeliveryTime``.
	//  For ``DeliveryTime``, the time in the log file names corresponds to the delivery time for the log files.
	//   For ``EventTime``, The logs delivered are for a specific day only. The year, month, and day correspond to the day on which the event occurred, and the hour, minutes and seconds are set to 00 in the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3_bucket#partition_date_source S3Bucket#partition_date_source}
	PartitionDateSource *string `field:"optional" json:"partitionDateSource" yaml:"partitionDateSource"`
}

