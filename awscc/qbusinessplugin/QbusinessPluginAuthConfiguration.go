// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessplugin


type QbusinessPluginAuthConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/qbusiness_plugin#basic_auth_configuration QbusinessPlugin#basic_auth_configuration}.
	BasicAuthConfiguration *QbusinessPluginAuthConfigurationBasicAuthConfiguration `field:"optional" json:"basicAuthConfiguration" yaml:"basicAuthConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/qbusiness_plugin#no_auth_configuration QbusinessPlugin#no_auth_configuration}.
	NoAuthConfiguration *string `field:"optional" json:"noAuthConfiguration" yaml:"noAuthConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/qbusiness_plugin#o_auth_2_client_credential_configuration QbusinessPlugin#o_auth_2_client_credential_configuration}.
	OAuth2ClientCredentialConfiguration *QbusinessPluginAuthConfigurationOAuth2ClientCredentialConfiguration `field:"optional" json:"oAuth2ClientCredentialConfiguration" yaml:"oAuth2ClientCredentialConfiguration"`
}

