// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseassetmodel


type IotsitewiseAssetModelEnforcedAssetModelInterfaceRelationshipsPropertyMappings struct {
	// The external ID of the enforced asset model property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_asset_model#asset_model_property_external_id IotsitewiseAssetModel#asset_model_property_external_id}
	AssetModelPropertyExternalId *string `field:"optional" json:"assetModelPropertyExternalId" yaml:"assetModelPropertyExternalId"`
	// The logical ID of the enforced asset model property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_asset_model#asset_model_property_logical_id IotsitewiseAssetModel#asset_model_property_logical_id}
	AssetModelPropertyLogicalId *string `field:"optional" json:"assetModelPropertyLogicalId" yaml:"assetModelPropertyLogicalId"`
	// The external ID of the enforced interface property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_asset_model#interface_asset_model_property_external_id IotsitewiseAssetModel#interface_asset_model_property_external_id}
	InterfaceAssetModelPropertyExternalId *string `field:"optional" json:"interfaceAssetModelPropertyExternalId" yaml:"interfaceAssetModelPropertyExternalId"`
}

