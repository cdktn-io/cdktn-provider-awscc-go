// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#content_type GlueConnectionType#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#jwt_token GlueConnectionType#jwt_token}
	JwtToken *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesJwtToken `field:"optional" json:"jwtToken" yaml:"jwtToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#request_method GlueConnectionType#request_method}.
	RequestMethod *string `field:"optional" json:"requestMethod" yaml:"requestMethod"`
	// Defines a property configuration for connection types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#token_url GlueConnectionType#token_url}
	TokenUrl *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrl `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection_type#token_url_parameters GlueConnectionType#token_url_parameters}.
	TokenUrlParameters interface{} `field:"optional" json:"tokenUrlParameters" yaml:"tokenUrlParameters"`
}

