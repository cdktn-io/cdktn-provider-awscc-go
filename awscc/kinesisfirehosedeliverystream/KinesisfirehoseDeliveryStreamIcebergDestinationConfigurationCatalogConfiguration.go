// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationCatalogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisfirehose_delivery_stream#catalog_arn KinesisfirehoseDeliveryStream#catalog_arn}.
	CatalogArn *string `field:"optional" json:"catalogArn" yaml:"catalogArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisfirehose_delivery_stream#warehouse_location KinesisfirehoseDeliveryStream#warehouse_location}.
	WarehouseLocation *string `field:"optional" json:"warehouseLocation" yaml:"warehouseLocation"`
}

