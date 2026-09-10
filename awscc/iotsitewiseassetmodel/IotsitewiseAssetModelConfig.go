// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseassetmodel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotsitewiseAssetModelConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A unique, friendly name for the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#asset_model_name IotsitewiseAssetModel#asset_model_name}
	AssetModelName *string `field:"required" json:"assetModelName" yaml:"assetModelName"`
	// The composite asset models that are part of this asset model.
	//
	// Composite asset models are asset models that contain specific properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#asset_model_composite_models IotsitewiseAssetModel#asset_model_composite_models}
	AssetModelCompositeModels interface{} `field:"optional" json:"assetModelCompositeModels" yaml:"assetModelCompositeModels"`
	// A description for the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#asset_model_description IotsitewiseAssetModel#asset_model_description}
	AssetModelDescription *string `field:"optional" json:"assetModelDescription" yaml:"assetModelDescription"`
	// The external ID of the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#asset_model_external_id IotsitewiseAssetModel#asset_model_external_id}
	AssetModelExternalId *string `field:"optional" json:"assetModelExternalId" yaml:"assetModelExternalId"`
	// The hierarchy definitions of the asset model.
	//
	// Each hierarchy specifies an asset model whose assets can be children of any other assets created from this asset model. You can specify up to 10 hierarchies per asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#asset_model_hierarchies IotsitewiseAssetModel#asset_model_hierarchies}
	AssetModelHierarchies interface{} `field:"optional" json:"assetModelHierarchies" yaml:"assetModelHierarchies"`
	// The property definitions of the asset model. You can specify up to 200 properties per asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#asset_model_properties IotsitewiseAssetModel#asset_model_properties}
	AssetModelProperties interface{} `field:"optional" json:"assetModelProperties" yaml:"assetModelProperties"`
	// The type of the asset model (ASSET_MODEL OR COMPONENT_MODEL or INTERFACE).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#asset_model_type IotsitewiseAssetModel#asset_model_type}
	AssetModelType *string `field:"optional" json:"assetModelType" yaml:"assetModelType"`
	// a list of asset model and interface relationships.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#enforced_asset_model_interface_relationships IotsitewiseAssetModel#enforced_asset_model_interface_relationships}
	EnforcedAssetModelInterfaceRelationships interface{} `field:"optional" json:"enforcedAssetModelInterfaceRelationships" yaml:"enforcedAssetModelInterfaceRelationships"`
	// A list of key-value pairs that contain metadata for the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#tags IotsitewiseAssetModel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

