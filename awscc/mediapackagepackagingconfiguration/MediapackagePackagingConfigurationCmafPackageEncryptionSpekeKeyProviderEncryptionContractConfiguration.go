// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagepackagingconfiguration


type MediapackagePackagingConfigurationCmafPackageEncryptionSpekeKeyProviderEncryptionContractConfiguration struct {
	// A collection of audio encryption presets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackage_packaging_configuration#preset_speke_20_audio MediapackagePackagingConfiguration#preset_speke_20_audio}
	PresetSpeke20Audio *string `field:"optional" json:"presetSpeke20Audio" yaml:"presetSpeke20Audio"`
	// A collection of video encryption presets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackage_packaging_configuration#preset_speke_20_video MediapackagePackagingConfiguration#preset_speke_20_video}
	PresetSpeke20Video *string `field:"optional" json:"presetSpeke20Video" yaml:"presetSpeke20Video"`
}

