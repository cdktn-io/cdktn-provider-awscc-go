// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection


type GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2Credentials struct {
	// The access token used in the authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_connection#access_token GlueConnection#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The JSON Web Token (JWT) used when the authentication type is OAuth2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_connection#jwt_token GlueConnection#jwt_token}
	JwtToken *string `field:"optional" json:"jwtToken" yaml:"jwtToken"`
	// The refresh token used when the authentication type is OAuth2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_connection#refresh_token GlueConnection#refresh_token}
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
	// The client application client secret if the client application is user managed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_connection#user_managed_client_application_client_secret GlueConnection#user_managed_client_application_client_secret}
	UserManagedClientApplicationClientSecret *string `field:"optional" json:"userManagedClientApplicationClientSecret" yaml:"userManagedClientApplicationClientSecret"`
}

