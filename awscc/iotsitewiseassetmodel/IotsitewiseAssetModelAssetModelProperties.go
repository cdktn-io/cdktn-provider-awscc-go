// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelProperties struct {
	// The data type of the asset model property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#data_type IotsitewiseAssetModel#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// The data type of the structure for this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#data_type_spec IotsitewiseAssetModel#data_type_spec}
	DataTypeSpec *string `field:"optional" json:"dataTypeSpec" yaml:"dataTypeSpec"`
	// The External ID of the Asset Model Property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#external_id IotsitewiseAssetModel#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The ID of the Asset Model Property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#id IotsitewiseAssetModel#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Customer provided Logical ID for property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#logical_id IotsitewiseAssetModel#logical_id}
	LogicalId *string `field:"optional" json:"logicalId" yaml:"logicalId"`
	// The name of the asset model property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#name IotsitewiseAssetModel#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The property type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#type IotsitewiseAssetModel#type}
	Type *IotsitewiseAssetModelAssetModelPropertiesType `field:"optional" json:"type" yaml:"type"`
	// The unit of the asset model property, such as Newtons or RPM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#unit IotsitewiseAssetModel#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

