// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketo struct {
	// The credentials used to access protected resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_connector_profile#access_token AppflowConnectorProfile#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The identi?er for the desired client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_connector_profile#client_id AppflowConnectorProfile#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client secret used by the oauth client to authenticate to the authorization server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_connector_profile#client_secret AppflowConnectorProfile#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The oauth needed to request security tokens from the connector endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_connector_profile#connector_o_auth_request AppflowConnectorProfile#connector_o_auth_request}
	ConnectorOAuthRequest *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoConnectorOAuthRequest `field:"optional" json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
}

