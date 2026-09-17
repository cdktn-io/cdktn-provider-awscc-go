// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitystoreuser


type IdentitystoreUserName struct {
	// The family name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#family_name IdentitystoreUser#family_name}
	FamilyName *string `field:"optional" json:"familyName" yaml:"familyName"`
	// A string containing a formatted version of the name for display.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#formatted IdentitystoreUser#formatted}
	Formatted *string `field:"optional" json:"formatted" yaml:"formatted"`
	// The given name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#given_name IdentitystoreUser#given_name}
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
	// The honorific prefix of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#honorific_prefix IdentitystoreUser#honorific_prefix}
	HonorificPrefix *string `field:"optional" json:"honorificPrefix" yaml:"honorificPrefix"`
	// The honorific suffix of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#honorific_suffix IdentitystoreUser#honorific_suffix}
	HonorificSuffix *string `field:"optional" json:"honorificSuffix" yaml:"honorificSuffix"`
	// The middle name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/identitystore_user#middle_name IdentitystoreUser#middle_name}
	MiddleName *string `field:"optional" json:"middleName" yaml:"middleName"`
}

