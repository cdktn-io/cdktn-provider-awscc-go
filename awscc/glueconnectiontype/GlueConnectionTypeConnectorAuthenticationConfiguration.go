// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeConnectorAuthenticationConfiguration struct {
	// A list of authentication types supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#authentication_types GlueConnectionType#authentication_types}
	AuthenticationTypes *[]*string `field:"optional" json:"authenticationTypes" yaml:"authenticationTypes"`
	// Basic authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#basic_authentication_properties GlueConnectionType#basic_authentication_properties}
	BasicAuthenticationProperties *GlueConnectionTypeConnectorAuthenticationConfigurationBasicAuthenticationProperties `field:"optional" json:"basicAuthenticationProperties" yaml:"basicAuthenticationProperties"`
	// Custom authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#custom_authentication_properties GlueConnectionType#custom_authentication_properties}
	CustomAuthenticationProperties *GlueConnectionTypeConnectorAuthenticationConfigurationCustomAuthenticationProperties `field:"optional" json:"customAuthenticationProperties" yaml:"customAuthenticationProperties"`
	// OAuth2 configuration container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#o_auth_2_properties GlueConnectionType#o_auth_2_properties}
	OAuth2Properties *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2Properties `field:"optional" json:"oAuth2Properties" yaml:"oAuth2Properties"`
}

