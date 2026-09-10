// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlackConnectorOAuthRequest struct {
	// The code provided by the connector when it has been authenticated via the connected app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appflow_connector_profile#auth_code AppflowConnectorProfile#auth_code}
	AuthCode *string `field:"optional" json:"authCode" yaml:"authCode"`
	// The URL to which the authentication server redirects the browser after authorization has been granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appflow_connector_profile#redirect_uri AppflowConnectorProfile#redirect_uri}
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

