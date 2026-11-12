// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateFilesCodeSigningCustomCodeSigningCertificateChain struct {
	// The name of the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#certificate_name IotOtaUpdate#certificate_name}
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// A base64 encoded binary representation of the code signing certificate chain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#inline_document IotOtaUpdate#inline_document}
	InlineDocument *string `field:"optional" json:"inlineDocument" yaml:"inlineDocument"`
}

