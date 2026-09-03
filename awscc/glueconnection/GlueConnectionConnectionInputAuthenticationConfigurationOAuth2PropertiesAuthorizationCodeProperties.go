// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection


type GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties struct {
	// The authorization code used in the authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#authorization_code GlueConnection#authorization_code}
	AuthorizationCode *string `field:"optional" json:"authorizationCode" yaml:"authorizationCode"`
	// The redirect URI where the user gets redirected to by authorization server when issuing an authorization code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#redirect_uri GlueConnection#redirect_uri}
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

