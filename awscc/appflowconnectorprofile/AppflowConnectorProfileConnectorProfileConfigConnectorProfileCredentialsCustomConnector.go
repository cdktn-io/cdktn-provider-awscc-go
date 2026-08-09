// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appflow_connector_profile#api_key AppflowConnectorProfile#api_key}.
	ApiKey *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appflow_connector_profile#authentication_type AppflowConnectorProfile#authentication_type}.
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appflow_connector_profile#basic AppflowConnectorProfile#basic}.
	Basic *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorBasic `field:"optional" json:"basic" yaml:"basic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appflow_connector_profile#custom AppflowConnectorProfile#custom}.
	Custom *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorCustom `field:"optional" json:"custom" yaml:"custom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appflow_connector_profile#oauth_2 AppflowConnectorProfile#oauth_2}.
	Oauth2 *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
}

