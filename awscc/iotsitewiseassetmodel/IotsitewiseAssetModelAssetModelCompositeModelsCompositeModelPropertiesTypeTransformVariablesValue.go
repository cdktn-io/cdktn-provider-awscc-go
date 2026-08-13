// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesTypeTransformVariablesValue struct {
	// The External ID of the hierarchy that is trying to be referenced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#hierarchy_external_id IotsitewiseAssetModel#hierarchy_external_id}
	HierarchyExternalId *string `field:"optional" json:"hierarchyExternalId" yaml:"hierarchyExternalId"`
	// The ID of the hierarchy that is trying to be referenced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#hierarchy_id IotsitewiseAssetModel#hierarchy_id}
	HierarchyId *string `field:"optional" json:"hierarchyId" yaml:"hierarchyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#hierarchy_logical_id IotsitewiseAssetModel#hierarchy_logical_id}.
	HierarchyLogicalId *string `field:"optional" json:"hierarchyLogicalId" yaml:"hierarchyLogicalId"`
	// The External ID of the property that is trying to be referenced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#property_external_id IotsitewiseAssetModel#property_external_id}
	PropertyExternalId *string `field:"optional" json:"propertyExternalId" yaml:"propertyExternalId"`
	// The ID of the property that is trying to be referenced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#property_id IotsitewiseAssetModel#property_id}
	PropertyId *string `field:"optional" json:"propertyId" yaml:"propertyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#property_logical_id IotsitewiseAssetModel#property_logical_id}.
	PropertyLogicalId *string `field:"optional" json:"propertyLogicalId" yaml:"propertyLogicalId"`
	// The path of the property that is trying to be referenced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#property_path IotsitewiseAssetModel#property_path}
	PropertyPath interface{} `field:"optional" json:"propertyPath" yaml:"propertyPath"`
}

