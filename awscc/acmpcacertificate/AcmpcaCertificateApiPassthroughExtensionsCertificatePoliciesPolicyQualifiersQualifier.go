// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersQualifier struct {
	// Contains a pointer to a certification practice statement (CPS) published by the CA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/acmpca_certificate#cps_uri AcmpcaCertificate#cps_uri}
	CpsUri *string `field:"optional" json:"cpsUri" yaml:"cpsUri"`
}

