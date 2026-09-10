// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelIcebergDestinationConfigurationDestinationTableListStruct struct {
	// The destination database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/msk_channel#destination_database_name MskChannel#destination_database_name}
	DestinationDatabaseName *string `field:"optional" json:"destinationDatabaseName" yaml:"destinationDatabaseName"`
	// The destination table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/msk_channel#destination_table_name MskChannel#destination_table_name}
	DestinationTableName *string `field:"optional" json:"destinationTableName" yaml:"destinationTableName"`
	// Partition specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/msk_channel#partition_spec MskChannel#partition_spec}
	PartitionSpec *MskChannelIcebergDestinationConfigurationDestinationTableListPartitionSpec `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
}

