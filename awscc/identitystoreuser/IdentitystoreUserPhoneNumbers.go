// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitystoreuser


type IdentitystoreUserPhoneNumbers struct {
	// Whether this is the primary phone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#primary IdentitystoreUser#primary}
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The type of phone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#type IdentitystoreUser#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The phone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#value IdentitystoreUser#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

