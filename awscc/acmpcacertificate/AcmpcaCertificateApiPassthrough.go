// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateApiPassthrough struct {
	// Specifies X.509 extension information for a certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/acmpca_certificate#extensions AcmpcaCertificate#extensions}
	Extensions *AcmpcaCertificateApiPassthroughExtensions `field:"optional" json:"extensions" yaml:"extensions"`
	// Contains information about the certificate subject.
	//
	// The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/acmpca_certificate#subject AcmpcaCertificate#subject}
	Subject *AcmpcaCertificateApiPassthroughSubject `field:"optional" json:"subject" yaml:"subject"`
}

