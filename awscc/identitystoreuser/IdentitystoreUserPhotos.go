// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitystoreuser


type IdentitystoreUserPhotos struct {
	// A display name for the photo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/identitystore_user#display IdentitystoreUser#display}
	Display *string `field:"optional" json:"display" yaml:"display"`
	// Whether this is the primary photo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/identitystore_user#primary IdentitystoreUser#primary}
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The type of photo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/identitystore_user#type IdentitystoreUser#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The photo data or URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/identitystore_user#value IdentitystoreUser#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

