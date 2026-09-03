// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetKinesisParameters struct {
	// The custom parameter used as the Kinesis partition key.
	//
	// For more information, see Amazon Kinesis Streams Key Concepts in the Amazon Kinesis Streams Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/scheduler_schedule#partition_key SchedulerSchedule#partition_key}
	PartitionKey *string `field:"optional" json:"partitionKey" yaml:"partitionKey"`
}

