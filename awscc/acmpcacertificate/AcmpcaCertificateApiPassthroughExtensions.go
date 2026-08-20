// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensions struct {
	// Contains a sequence of one or more policy information terms, each of which consists of an object identifier (OID) and optional qualifiers.
	//
	// For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier).
	//  In an end-entity certificate, these terms indicate the policy under which the certificate was issued and the purposes for which it may be used. In a CA certificate, these terms limit the set of policies for certification paths that include this certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#certificate_policies AcmpcaCertificate#certificate_policies}
	CertificatePolicies interface{} `field:"optional" json:"certificatePolicies" yaml:"certificatePolicies"`
	// Contains a sequence of one or more X.509 extensions, each of which consists of an object identifier (OID), a base64-encoded value, and the critical flag. For more information, see the [Global OID reference database.](https://docs.aws.amazon.com/https://oidref.com/2.5.29).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#custom_extensions AcmpcaCertificate#custom_extensions}
	CustomExtensions interface{} `field:"optional" json:"customExtensions" yaml:"customExtensions"`
	// Specifies additional purposes for which the certified public key may be used other than basic purposes indicated in the ``KeyUsage`` extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#extended_key_usage AcmpcaCertificate#extended_key_usage}
	ExtendedKeyUsage interface{} `field:"optional" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// Defines one or more purposes for which the key contained in the certificate can be used.
	//
	// Default value for each option is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#key_usage AcmpcaCertificate#key_usage}
	KeyUsage *AcmpcaCertificateApiPassthroughExtensionsKeyUsage `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// The subject alternative name extension allows identities to be bound to the subject of the certificate.
	//
	// These identities may be included in addition to or in place of the identity in the subject field of the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#subject_alternative_names AcmpcaCertificate#subject_alternative_names}
	SubjectAlternativeNames interface{} `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

