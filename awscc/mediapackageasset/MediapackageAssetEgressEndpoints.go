// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackageasset


type MediapackageAssetEgressEndpoints struct {
	// The ID of the PackagingConfiguration being applied to the Asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackage_asset#packaging_configuration_id MediapackageAsset#packaging_configuration_id}
	PackagingConfigurationId *string `field:"optional" json:"packagingConfigurationId" yaml:"packagingConfigurationId"`
	// The URL of the parent manifest for the repackaged Asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackage_asset#url MediapackageAsset#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

