// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV2PrivateKeyFlags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#client_version PcaconnectoradTemplate#client_version}.
	ClientVersion *string `field:"optional" json:"clientVersion" yaml:"clientVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#exportable_key PcaconnectoradTemplate#exportable_key}.
	ExportableKey interface{} `field:"optional" json:"exportableKey" yaml:"exportableKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcaconnectorad_template#strong_key_protection_required PcaconnectoradTemplate#strong_key_protection_required}.
	StrongKeyProtectionRequired interface{} `field:"optional" json:"strongKeyProtectionRequired" yaml:"strongKeyProtectionRequired"`
}

