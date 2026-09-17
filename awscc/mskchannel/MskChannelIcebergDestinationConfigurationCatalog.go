// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelIcebergDestinationConfigurationCatalog struct {
	// The ARN of the catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#catalog_arn MskChannel#catalog_arn}
	CatalogArn *string `field:"optional" json:"catalogArn" yaml:"catalogArn"`
	// The warehouse location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#warehouse_location MskChannel#warehouse_location}
	WarehouseLocation *string `field:"optional" json:"warehouseLocation" yaml:"warehouseLocation"`
}

