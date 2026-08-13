// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommand struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#copy_options KinesisfirehoseDeliveryStream#copy_options}.
	CopyOptions *string `field:"optional" json:"copyOptions" yaml:"copyOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#data_table_columns KinesisfirehoseDeliveryStream#data_table_columns}.
	DataTableColumns *string `field:"optional" json:"dataTableColumns" yaml:"dataTableColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#data_table_name KinesisfirehoseDeliveryStream#data_table_name}.
	DataTableName *string `field:"optional" json:"dataTableName" yaml:"dataTableName"`
}

