// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsMcpServerAuthorizationConfig struct {
	// API key authentication details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/devopsagent_service#api_key DevopsagentService#api_key}
	ApiKey *DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Bearer token authentication details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/devopsagent_service#bearer_token DevopsagentService#bearer_token}
	BearerToken *DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigBearerToken `field:"optional" json:"bearerToken" yaml:"bearerToken"`
	// MCP server OAuth client credentials configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/devopsagent_service#o_auth_client_credentials DevopsagentService#o_auth_client_credentials}
	OAuthClientCredentials *DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigOAuthClientCredentials `field:"optional" json:"oAuthClientCredentials" yaml:"oAuthClientCredentials"`
}

