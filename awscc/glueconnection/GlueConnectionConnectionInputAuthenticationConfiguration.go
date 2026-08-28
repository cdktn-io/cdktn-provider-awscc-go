// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection


type GlueConnectionConnectionInputAuthenticationConfiguration struct {
	// A structure containing the authentication configuration in the CreateConnection request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_connection#authentication_type GlueConnection#authentication_type}
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// For supplying basic auth credentials when not providing a SecretArn value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_connection#basic_authentication_credentials GlueConnection#basic_authentication_credentials}
	BasicAuthenticationCredentials *GlueConnectionConnectionInputAuthenticationConfigurationBasicAuthenticationCredentials `field:"optional" json:"basicAuthenticationCredentials" yaml:"basicAuthenticationCredentials"`
	// A structure containing the authentication credentials in the CreateConnection request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_connection#custom_authentication_credentials GlueConnection#custom_authentication_credentials}
	CustomAuthenticationCredentials *string `field:"optional" json:"customAuthenticationCredentials" yaml:"customAuthenticationCredentials"`
	// The Amazon Resource Name (ARN) of the KMS key used in the authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_connection#kms_key_arn GlueConnection#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// A structure containing properties for OAuth2 in the CreateConnection request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_connection#o_auth_2_properties GlueConnection#o_auth_2_properties}
	OAuth2Properties *GlueConnectionConnectionInputAuthenticationConfigurationOAuth2Properties `field:"optional" json:"oAuth2Properties" yaml:"oAuth2Properties"`
	// The secret manager ARN to store credentials in the CreateConnection request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_connection#secret_arn GlueConnection#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

