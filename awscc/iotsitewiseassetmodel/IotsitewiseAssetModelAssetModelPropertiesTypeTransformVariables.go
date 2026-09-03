// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariables struct {
	// The friendly name of the variable to be used in the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_asset_model#name IotsitewiseAssetModel#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The variable that identifies an asset property from which to use values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotsitewise_asset_model#value IotsitewiseAssetModel#value}
	Value *IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValue `field:"optional" json:"value" yaml:"value"`
}

