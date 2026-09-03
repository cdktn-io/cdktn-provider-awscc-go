// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelHierarchies struct {
	// The ID of the asset model.
	//
	// All assets in this hierarchy must be instances of the child AssetModelId asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_asset_model#child_asset_model_id IotsitewiseAssetModel#child_asset_model_id}
	ChildAssetModelId *string `field:"optional" json:"childAssetModelId" yaml:"childAssetModelId"`
	// Customer provided external ID for hierarchy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_asset_model#external_id IotsitewiseAssetModel#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Customer provided actual ID for hierarchy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_asset_model#id IotsitewiseAssetModel#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Customer provided logical ID for hierarchy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_asset_model#logical_id IotsitewiseAssetModel#logical_id}
	LogicalId *string `field:"optional" json:"logicalId" yaml:"logicalId"`
	// The name of the asset model hierarchy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_asset_model#name IotsitewiseAssetModel#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

