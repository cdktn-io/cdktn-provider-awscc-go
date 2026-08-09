// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateApiPassthroughSubjectCustomAttributes struct {
	// Specifies the object identifier (OID) of the attribute type of the relative distinguished name (RDN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/acmpca_certificate#object_identifier AcmpcaCertificate#object_identifier}
	ObjectIdentifier *string `field:"optional" json:"objectIdentifier" yaml:"objectIdentifier"`
	// Specifies the attribute value of relative distinguished name (RDN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/acmpca_certificate#value AcmpcaCertificate#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

