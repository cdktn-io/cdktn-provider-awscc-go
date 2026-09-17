// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesClientCredentialsProperties struct {
	// Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#client_id GlueConnectionType#client_id}
	ClientId *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesClientCredentialsPropertiesClientId `field:"optional" json:"clientId" yaml:"clientId"`
	// Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#client_secret GlueConnectionType#client_secret}
	ClientSecret *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesClientCredentialsPropertiesClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#content_type GlueConnectionType#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#request_method GlueConnectionType#request_method}.
	RequestMethod *string `field:"optional" json:"requestMethod" yaml:"requestMethod"`
	// Defines a property configuration for connection types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#scope GlueConnectionType#scope}
	Scope *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesClientCredentialsPropertiesScope `field:"optional" json:"scope" yaml:"scope"`
	// Defines a property configuration for connection types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#token_url GlueConnectionType#token_url}
	TokenUrl *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesClientCredentialsPropertiesTokenUrl `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#token_url_parameters GlueConnectionType#token_url_parameters}.
	TokenUrlParameters interface{} `field:"optional" json:"tokenUrlParameters" yaml:"tokenUrlParameters"`
}

