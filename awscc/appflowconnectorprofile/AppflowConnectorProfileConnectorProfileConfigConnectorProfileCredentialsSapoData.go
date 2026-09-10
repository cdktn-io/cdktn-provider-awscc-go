// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appflow_connector_profile#basic_auth_credentials AppflowConnectorProfile#basic_auth_credentials}.
	BasicAuthCredentials *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataBasicAuthCredentials `field:"optional" json:"basicAuthCredentials" yaml:"basicAuthCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appflow_connector_profile#o_auth_credentials AppflowConnectorProfile#o_auth_credentials}.
	OAuthCredentials *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataOAuthCredentials `field:"optional" json:"oAuthCredentials" yaml:"oAuthCredentials"`
}

