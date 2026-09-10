// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsPagerDutyAuthorizationConfig struct {
	// OAuth client credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_service#o_auth_client_credentials DevopsagentService#o_auth_client_credentials}
	OAuthClientCredentials *DevopsagentServiceServiceDetailsPagerDutyAuthorizationConfigOAuthClientCredentials `field:"optional" json:"oAuthClientCredentials" yaml:"oAuthClientCredentials"`
}

