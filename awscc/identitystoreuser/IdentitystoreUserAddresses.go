// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitystoreuser


type IdentitystoreUserAddresses struct {
	// The country of the address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/identitystore_user#country IdentitystoreUser#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// A formatted version of the address for display.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/identitystore_user#formatted IdentitystoreUser#formatted}
	Formatted *string `field:"optional" json:"formatted" yaml:"formatted"`
	// A string of the address locality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/identitystore_user#locality IdentitystoreUser#locality}
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// The postal code of the address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/identitystore_user#postal_code IdentitystoreUser#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Whether this is the primary address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/identitystore_user#primary IdentitystoreUser#primary}
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The region of the address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/identitystore_user#region IdentitystoreUser#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The street of the address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/identitystore_user#street_address IdentitystoreUser#street_address}
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// The type of address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/identitystore_user#type IdentitystoreUser#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

