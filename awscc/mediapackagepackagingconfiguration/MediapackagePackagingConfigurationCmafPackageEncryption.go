// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagepackagingconfiguration


type MediapackagePackagingConfigurationCmafPackageEncryption struct {
	// A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediapackage_packaging_configuration#speke_key_provider MediapackagePackagingConfiguration#speke_key_provider}
	SpekeKeyProvider *MediapackagePackagingConfigurationCmafPackageEncryptionSpekeKeyProvider `field:"optional" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
}

