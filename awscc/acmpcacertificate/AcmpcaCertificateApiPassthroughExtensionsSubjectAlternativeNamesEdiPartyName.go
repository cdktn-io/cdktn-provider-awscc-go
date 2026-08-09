// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesEdiPartyName struct {
	// Specifies the name assigner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/acmpca_certificate#name_assigner AcmpcaCertificate#name_assigner}
	NameAssigner *string `field:"optional" json:"nameAssigner" yaml:"nameAssigner"`
	// Specifies the party name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/acmpca_certificate#party_name AcmpcaCertificate#party_name}
	PartyName *string `field:"optional" json:"partyName" yaml:"partyName"`
}

