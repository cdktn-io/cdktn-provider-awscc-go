// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelIcebergDestinationConfigurationDestinationTableListPartitionSpec struct {
	// Partition strategy for MSK channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/msk_channel#partition_strategy MskChannel#partition_strategy}
	PartitionStrategy *string `field:"optional" json:"partitionStrategy" yaml:"partitionStrategy"`
	// Source list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/msk_channel#source_list MskChannel#source_list}
	SourceList interface{} `field:"optional" json:"sourceList" yaml:"sourceList"`
}

