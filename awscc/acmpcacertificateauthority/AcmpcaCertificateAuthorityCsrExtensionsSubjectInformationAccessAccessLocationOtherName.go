// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority


type AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOtherName struct {
	// String that contains X.509 ObjectIdentifier information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/acmpca_certificate_authority#type_id AcmpcaCertificateAuthority#type_id}
	TypeId *string `field:"optional" json:"typeId" yaml:"typeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/acmpca_certificate_authority#value AcmpcaCertificateAuthority#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

