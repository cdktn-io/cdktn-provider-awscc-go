// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel


type KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListPartitionSpecPartitionFields struct {
	// The name of the source column on which the transform is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesis_channel#source_name KinesisChannel#source_name}
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
	// The partitioning transform applied to the SourceName column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesis_channel#transform KinesisChannel#transform}
	Transform *string `field:"optional" json:"transform" yaml:"transform"`
}

