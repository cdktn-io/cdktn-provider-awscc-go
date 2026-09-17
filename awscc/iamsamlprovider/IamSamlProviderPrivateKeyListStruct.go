// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamsamlprovider


type IamSamlProviderPrivateKeyListStruct struct {
	// The unique identifier for the SAML private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iam_saml_provider#key_id IamSamlProvider#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// The date and time, in <a href=\"http://www.iso.org/iso/iso8601\">ISO 8601 date-time </a> format, when the private key was uploaded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iam_saml_provider#timestamp IamSamlProvider#timestamp}
	Timestamp *string `field:"optional" json:"timestamp" yaml:"timestamp"`
}

