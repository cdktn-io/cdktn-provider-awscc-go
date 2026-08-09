// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiers struct {
	// Identifies the qualifier modifying a ``CertPolicyId``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/acmpca_certificate#policy_qualifier_id AcmpcaCertificate#policy_qualifier_id}
	PolicyQualifierId *string `field:"optional" json:"policyQualifierId" yaml:"policyQualifierId"`
	// Defines the qualifier type.
	//
	// AWS Private CA supports the use of a URI for a CPS qualifier in this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/acmpca_certificate#qualifier AcmpcaCertificate#qualifier}
	Qualifier *AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersQualifier `field:"optional" json:"qualifier" yaml:"qualifier"`
}

