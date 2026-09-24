// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforce struct {
	// The credentials used to access protected resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appflow_connector_profile#access_token AppflowConnectorProfile#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The client credentials to fetch access token and refresh token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appflow_connector_profile#client_credentials_arn AppflowConnectorProfile#client_credentials_arn}
	ClientCredentialsArn *string `field:"optional" json:"clientCredentialsArn" yaml:"clientCredentialsArn"`
	// The oauth needed to request security tokens from the connector endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appflow_connector_profile#connector_o_auth_request AppflowConnectorProfile#connector_o_auth_request}
	ConnectorOAuthRequest *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforceConnectorOAuthRequest `field:"optional" json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
	// The credentials used to acquire new access tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appflow_connector_profile#refresh_token AppflowConnectorProfile#refresh_token}
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

