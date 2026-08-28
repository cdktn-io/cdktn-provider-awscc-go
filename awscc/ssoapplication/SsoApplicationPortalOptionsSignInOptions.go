// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssoapplication


type SsoApplicationPortalOptionsSignInOptions struct {
	// The URL that accepts authentication requests for an application, this is a required parameter if the Origin parameter is APPLICATION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sso_application#application_url SsoApplication#application_url}
	ApplicationUrl *string `field:"optional" json:"applicationUrl" yaml:"applicationUrl"`
	// This determines how IAM Identity Center navigates the user to the target application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sso_application#origin SsoApplication#origin}
	Origin *string `field:"optional" json:"origin" yaml:"origin"`
}

