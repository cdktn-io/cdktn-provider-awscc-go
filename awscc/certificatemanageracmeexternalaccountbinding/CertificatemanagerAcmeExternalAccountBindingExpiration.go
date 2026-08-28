// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certificatemanageracmeexternalaccountbinding


type CertificatemanagerAcmeExternalAccountBindingExpiration struct {
	// The time unit for the expiration value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/certificatemanager_acme_external_account_binding#type CertificatemanagerAcmeExternalAccountBinding#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The expiration value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/certificatemanager_acme_external_account_binding#value CertificatemanagerAcmeExternalAccountBinding#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

