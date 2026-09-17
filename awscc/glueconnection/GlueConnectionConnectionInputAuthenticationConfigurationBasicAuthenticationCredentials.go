// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection


type GlueConnectionConnectionInputAuthenticationConfigurationBasicAuthenticationCredentials struct {
	// The password used in the authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection#password GlueConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The username used in the authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_connection#username GlueConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

