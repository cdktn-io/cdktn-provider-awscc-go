// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package locationapikey


type LocationApiKeyRestrictionsAllowAndroidApps struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/location_api_key#certificate_fingerprint LocationApiKey#certificate_fingerprint}.
	CertificateFingerprint *string `field:"optional" json:"certificateFingerprint" yaml:"certificateFingerprint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/location_api_key#package LocationApiKey#package}.
	Package *string `field:"optional" json:"package" yaml:"package"`
}

