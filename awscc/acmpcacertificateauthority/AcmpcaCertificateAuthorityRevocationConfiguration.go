// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority


type AcmpcaCertificateAuthorityRevocationConfiguration struct {
	// Your certificate authority can create and maintain a certificate revocation list (CRL).
	//
	// A CRL contains information about certificates that have been revoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#crl_configuration AcmpcaCertificateAuthority#crl_configuration}
	CrlConfiguration *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration `field:"optional" json:"crlConfiguration" yaml:"crlConfiguration"`
	// Helps to configure online certificate status protocol (OCSP) responder for your certificate authority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#ocsp_configuration AcmpcaCertificateAuthority#ocsp_configuration}
	OcspConfiguration *AcmpcaCertificateAuthorityRevocationConfigurationOcspConfiguration `field:"optional" json:"ocspConfiguration" yaml:"ocspConfiguration"`
}

