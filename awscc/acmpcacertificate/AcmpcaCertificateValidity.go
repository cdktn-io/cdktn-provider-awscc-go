// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateValidity struct {
	// Specifies whether the ``Value`` parameter represents days, months, or years.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/acmpca_certificate#type AcmpcaCertificate#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A long integer interpreted according to the value of ``Type``, below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/acmpca_certificate#value AcmpcaCertificate#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

