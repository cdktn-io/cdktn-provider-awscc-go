// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeConnectorAuthenticationConfigurationBasicAuthenticationProperties struct {
	// Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#password GlueConnectionType#password}
	Password *GlueConnectionTypeConnectorAuthenticationConfigurationBasicAuthenticationPropertiesPassword `field:"optional" json:"password" yaml:"password"`
	// Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#username GlueConnectionType#username}
	Username *GlueConnectionTypeConnectorAuthenticationConfigurationBasicAuthenticationPropertiesUsername `field:"optional" json:"username" yaml:"username"`
}

