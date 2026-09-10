// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsKeyUsage struct {
	// Key can be used to sign CRLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/acmpca_certificate#crl_sign AcmpcaCertificate#crl_sign}
	CrlSign interface{} `field:"optional" json:"crlSign" yaml:"crlSign"`
	// Key can be used to decipher data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/acmpca_certificate#data_encipherment AcmpcaCertificate#data_encipherment}
	DataEncipherment interface{} `field:"optional" json:"dataEncipherment" yaml:"dataEncipherment"`
	// Key can be used only to decipher data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/acmpca_certificate#decipher_only AcmpcaCertificate#decipher_only}
	DecipherOnly interface{} `field:"optional" json:"decipherOnly" yaml:"decipherOnly"`
	// Key can be used for digital signing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/acmpca_certificate#digital_signature AcmpcaCertificate#digital_signature}
	DigitalSignature interface{} `field:"optional" json:"digitalSignature" yaml:"digitalSignature"`
	// Key can be used only to encipher data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/acmpca_certificate#encipher_only AcmpcaCertificate#encipher_only}
	EncipherOnly interface{} `field:"optional" json:"encipherOnly" yaml:"encipherOnly"`
	// Key can be used in a key-agreement protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/acmpca_certificate#key_agreement AcmpcaCertificate#key_agreement}
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// Key can be used to sign certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/acmpca_certificate#key_cert_sign AcmpcaCertificate#key_cert_sign}
	KeyCertSign interface{} `field:"optional" json:"keyCertSign" yaml:"keyCertSign"`
	// Key can be used to encipher data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/acmpca_certificate#key_encipherment AcmpcaCertificate#key_encipherment}
	KeyEncipherment interface{} `field:"optional" json:"keyEncipherment" yaml:"keyEncipherment"`
	// Key can be used for non-repudiation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/acmpca_certificate#non_repudiation AcmpcaCertificate#non_repudiation}
	NonRepudiation interface{} `field:"optional" json:"nonRepudiation" yaml:"nonRepudiation"`
}

