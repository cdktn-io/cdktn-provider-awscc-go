// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseasset


type IotsitewiseAssetAssetHierarchies struct {
	// The ID of the child asset to be associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_asset#child_asset_id IotsitewiseAsset#child_asset_id}
	ChildAssetId *string `field:"optional" json:"childAssetId" yaml:"childAssetId"`
	// String-friendly customer provided external ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_asset#external_id IotsitewiseAsset#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Customer provided actual UUID for property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_asset#id IotsitewiseAsset#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The LogicalID of a hierarchy in the parent asset's model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_asset#logical_id IotsitewiseAsset#logical_id}
	LogicalId *string `field:"optional" json:"logicalId" yaml:"logicalId"`
}

