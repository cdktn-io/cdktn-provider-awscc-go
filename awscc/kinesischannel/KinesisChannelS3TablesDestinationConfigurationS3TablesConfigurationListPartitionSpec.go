// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel


type KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListPartitionSpec struct {
	// List of partition fields that define how records are partitioned when written to the destination table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesis_channel#partition_fields KinesisChannel#partition_fields}
	PartitionFields interface{} `field:"optional" json:"partitionFields" yaml:"partitionFields"`
}

