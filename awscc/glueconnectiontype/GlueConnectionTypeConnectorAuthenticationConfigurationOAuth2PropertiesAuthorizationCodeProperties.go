// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties struct {
	// Defines a property configuration for connection types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#authorization_code GlueConnectionType#authorization_code}
	AuthorizationCode *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCode `field:"optional" json:"authorizationCode" yaml:"authorizationCode"`
	// Defines a property configuration for connection types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#authorization_code_url GlueConnectionType#authorization_code_url}
	AuthorizationCodeUrl *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeUrl `field:"optional" json:"authorizationCodeUrl" yaml:"authorizationCodeUrl"`
	// Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#client_id GlueConnectionType#client_id}
	ClientId *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientId `field:"optional" json:"clientId" yaml:"clientId"`
	// Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#client_secret GlueConnectionType#client_secret}
	ClientSecret *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#content_type GlueConnectionType#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Defines a property configuration for connection types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#prompt GlueConnectionType#prompt}
	Prompt *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesPrompt `field:"optional" json:"prompt" yaml:"prompt"`
	// Defines a property configuration for connection types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#redirect_uri GlueConnectionType#redirect_uri}
	RedirectUri *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesRedirectUri `field:"optional" json:"redirectUri" yaml:"redirectUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#request_method GlueConnectionType#request_method}.
	RequestMethod *string `field:"optional" json:"requestMethod" yaml:"requestMethod"`
	// Defines a property configuration for connection types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#scope GlueConnectionType#scope}
	Scope *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesScope `field:"optional" json:"scope" yaml:"scope"`
	// Defines a property configuration for connection types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#token_url GlueConnectionType#token_url}
	TokenUrl *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrl `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#token_url_parameters GlueConnectionType#token_url_parameters}.
	TokenUrlParameters interface{} `field:"optional" json:"tokenUrlParameters" yaml:"tokenUrlParameters"`
}

