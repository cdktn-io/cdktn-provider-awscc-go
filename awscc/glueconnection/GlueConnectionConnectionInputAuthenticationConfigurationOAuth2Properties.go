// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection


type GlueConnectionConnectionInputAuthenticationConfigurationOAuth2Properties struct {
	// The set of properties required for the the OAuth2 AUTHORIZATION_CODE grant type workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection#authorization_code_properties GlueConnection#authorization_code_properties}
	AuthorizationCodeProperties *GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties `field:"optional" json:"authorizationCodeProperties" yaml:"authorizationCodeProperties"`
	// The OAuth2 client app used for the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection#o_auth_2_client_application GlueConnection#o_auth_2_client_application}
	OAuth2ClientApplication *GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2ClientApplication `field:"optional" json:"oAuth2ClientApplication" yaml:"oAuth2ClientApplication"`
	// A structure containing the OAuth2 credentials used in the authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection#o_auth_2_credentials GlueConnection#o_auth_2_credentials}
	OAuth2Credentials *GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2Credentials `field:"optional" json:"oAuth2Credentials" yaml:"oAuth2Credentials"`
	// The grant type used in the authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection#o_auth_2_grant_type GlueConnection#o_auth_2_grant_type}
	OAuth2GrantType *string `field:"optional" json:"oAuth2GrantType" yaml:"oAuth2GrantType"`
	// The URL used in the authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection#token_url GlueConnection#token_url}
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// A map of key-value pairs used in the authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection#token_url_parameters_map GlueConnection#token_url_parameters_map}
	TokenUrlParametersMap *string `field:"optional" json:"tokenUrlParametersMap" yaml:"tokenUrlParametersMap"`
}

