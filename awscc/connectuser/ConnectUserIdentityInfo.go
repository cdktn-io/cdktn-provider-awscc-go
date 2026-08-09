// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuser


type ConnectUserIdentityInfo struct {
	// The email address. If you are using SAML for identity management and include this parameter, an error is returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_user#email ConnectUser#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// The first name. This is required if you are using Amazon Connect or SAML for identity management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_user#first_name ConnectUser#first_name}
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// The last name. This is required if you are using Amazon Connect or SAML for identity management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_user#last_name ConnectUser#last_name}
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// The mobile phone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_user#mobile ConnectUser#mobile}
	Mobile *string `field:"optional" json:"mobile" yaml:"mobile"`
	// The secondary email address.
	//
	// If you provide a secondary email, the user receives email notifications -- other than password reset notifications -- to this email address instead of to their primary email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_user#secondary_email ConnectUser#secondary_email}
	SecondaryEmail *string `field:"optional" json:"secondaryEmail" yaml:"secondaryEmail"`
}

