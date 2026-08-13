// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#attribute IotsitewiseAssetModel#attribute}.
	Attribute *IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesTypeAttribute `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#metric IotsitewiseAssetModel#metric}.
	Metric *IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesTypeMetric `field:"optional" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#transform IotsitewiseAssetModel#transform}.
	Transform *IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesTypeTransform `field:"optional" json:"transform" yaml:"transform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_asset_model#type_name IotsitewiseAssetModel#type_name}.
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
}

