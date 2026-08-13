// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsCertificatePolicies struct {
	// Specifies the object identifier (OID) of the certificate policy under which the certificate was issued.
	//
	// For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/acmpca_certificate#cert_policy_id AcmpcaCertificate#cert_policy_id}
	CertPolicyId *string `field:"optional" json:"certPolicyId" yaml:"certPolicyId"`
	// Modifies the given ``CertPolicyId`` with a qualifier. AWS Private CA supports the certification practice statement (CPS) qualifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/acmpca_certificate#policy_qualifiers AcmpcaCertificate#policy_qualifiers}
	PolicyQualifiers interface{} `field:"optional" json:"policyQualifiers" yaml:"policyQualifiers"`
}

