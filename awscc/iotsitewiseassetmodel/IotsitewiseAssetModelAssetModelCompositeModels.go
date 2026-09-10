// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelCompositeModels struct {
	// The component model ID for which the composite model is composed of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#composed_asset_model_id IotsitewiseAssetModel#composed_asset_model_id}
	ComposedAssetModelId *string `field:"optional" json:"composedAssetModelId" yaml:"composedAssetModelId"`
	// The property definitions of the asset model. You can specify up to 200 properties per asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#composite_model_properties IotsitewiseAssetModel#composite_model_properties}
	CompositeModelProperties interface{} `field:"optional" json:"compositeModelProperties" yaml:"compositeModelProperties"`
	// A description for the asset composite model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#description IotsitewiseAssetModel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The External ID of the composite model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#external_id IotsitewiseAssetModel#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The Actual ID of the composite model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#id IotsitewiseAssetModel#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A unique, friendly name for the asset composite model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#name IotsitewiseAssetModel#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parent composite model External ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#parent_asset_model_composite_model_external_id IotsitewiseAssetModel#parent_asset_model_composite_model_external_id}
	ParentAssetModelCompositeModelExternalId *string `field:"optional" json:"parentAssetModelCompositeModelExternalId" yaml:"parentAssetModelCompositeModelExternalId"`
	// The path of the composite model. This is only for derived composite models.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#path IotsitewiseAssetModel#path}
	Path *[]*string `field:"optional" json:"path" yaml:"path"`
	// The type of the composite model. For alarm composite models, this type is AWS/ALARM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#type IotsitewiseAssetModel#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

