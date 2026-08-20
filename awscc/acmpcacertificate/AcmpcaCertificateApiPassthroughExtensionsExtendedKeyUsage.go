// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsExtendedKeyUsage struct {
	// Specifies a custom ``ExtendedKeyUsage`` with an object identifier (OID).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#extended_key_usage_object_identifier AcmpcaCertificate#extended_key_usage_object_identifier}
	ExtendedKeyUsageObjectIdentifier *string `field:"optional" json:"extendedKeyUsageObjectIdentifier" yaml:"extendedKeyUsageObjectIdentifier"`
	// Specifies a standard ``ExtendedKeyUsage`` as defined as in [RFC 5280](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate#extended_key_usage_type AcmpcaCertificate#extended_key_usage_type}
	ExtendedKeyUsageType *string `field:"optional" json:"extendedKeyUsageType" yaml:"extendedKeyUsageType"`
}

