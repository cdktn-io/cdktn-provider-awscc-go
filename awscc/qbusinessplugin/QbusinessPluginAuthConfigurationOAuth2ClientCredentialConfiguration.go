// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessplugin


type QbusinessPluginAuthConfigurationOAuth2ClientCredentialConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/qbusiness_plugin#authorization_url QbusinessPlugin#authorization_url}.
	AuthorizationUrl *string `field:"optional" json:"authorizationUrl" yaml:"authorizationUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/qbusiness_plugin#role_arn QbusinessPlugin#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/qbusiness_plugin#secret_arn QbusinessPlugin#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/qbusiness_plugin#token_url QbusinessPlugin#token_url}.
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
}

