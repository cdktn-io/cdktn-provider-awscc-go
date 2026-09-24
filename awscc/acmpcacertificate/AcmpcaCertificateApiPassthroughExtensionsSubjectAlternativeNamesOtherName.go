// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOtherName struct {
	// Specifies an OID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/acmpca_certificate#type_id AcmpcaCertificate#type_id}
	TypeId *string `field:"optional" json:"typeId" yaml:"typeId"`
	// Specifies an OID value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/acmpca_certificate#value AcmpcaCertificate#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

