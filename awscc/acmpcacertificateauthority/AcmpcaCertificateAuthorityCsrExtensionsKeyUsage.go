// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority


type AcmpcaCertificateAuthorityCsrExtensionsKeyUsage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#crl_sign AcmpcaCertificateAuthority#crl_sign}.
	CrlSign interface{} `field:"optional" json:"crlSign" yaml:"crlSign"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#data_encipherment AcmpcaCertificateAuthority#data_encipherment}.
	DataEncipherment interface{} `field:"optional" json:"dataEncipherment" yaml:"dataEncipherment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#decipher_only AcmpcaCertificateAuthority#decipher_only}.
	DecipherOnly interface{} `field:"optional" json:"decipherOnly" yaml:"decipherOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#digital_signature AcmpcaCertificateAuthority#digital_signature}.
	DigitalSignature interface{} `field:"optional" json:"digitalSignature" yaml:"digitalSignature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#encipher_only AcmpcaCertificateAuthority#encipher_only}.
	EncipherOnly interface{} `field:"optional" json:"encipherOnly" yaml:"encipherOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#key_agreement AcmpcaCertificateAuthority#key_agreement}.
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#key_cert_sign AcmpcaCertificateAuthority#key_cert_sign}.
	KeyCertSign interface{} `field:"optional" json:"keyCertSign" yaml:"keyCertSign"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#key_encipherment AcmpcaCertificateAuthority#key_encipherment}.
	KeyEncipherment interface{} `field:"optional" json:"keyEncipherment" yaml:"keyEncipherment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#non_repudiation AcmpcaCertificateAuthority#non_repudiation}.
	NonRepudiation interface{} `field:"optional" json:"nonRepudiation" yaml:"nonRepudiation"`
}

