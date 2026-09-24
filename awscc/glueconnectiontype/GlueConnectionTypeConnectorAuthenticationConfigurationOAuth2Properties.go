// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2Properties struct {
	// OAuth2 authorization code configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#authorization_code_properties GlueConnectionType#authorization_code_properties}
	AuthorizationCodeProperties *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties `field:"optional" json:"authorizationCodeProperties" yaml:"authorizationCodeProperties"`
	// OAuth2 client credentials configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#client_credentials_properties GlueConnectionType#client_credentials_properties}
	ClientCredentialsProperties *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesClientCredentialsProperties `field:"optional" json:"clientCredentialsProperties" yaml:"clientCredentialsProperties"`
	// JWT bearer token configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#jwt_bearer_properties GlueConnectionType#jwt_bearer_properties}
	JwtBearerProperties *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerProperties `field:"optional" json:"jwtBearerProperties" yaml:"jwtBearerProperties"`
	// The OAuth2 grant type to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#o_auth_2_grant_type GlueConnectionType#o_auth_2_grant_type}
	OAuth2GrantType *string `field:"optional" json:"oAuth2GrantType" yaml:"oAuth2GrantType"`
}

