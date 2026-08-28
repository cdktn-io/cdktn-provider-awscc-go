// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagepackagingconfiguration


type MediapackagePackagingConfigurationMssPackageMssManifests struct {
	// An optional string to include in the name of the manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackage_packaging_configuration#manifest_name MediapackagePackagingConfiguration#manifest_name}
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// A StreamSelection configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackage_packaging_configuration#stream_selection MediapackagePackagingConfiguration#stream_selection}
	StreamSelection *MediapackagePackagingConfigurationMssPackageMssManifestsStreamSelection `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}

