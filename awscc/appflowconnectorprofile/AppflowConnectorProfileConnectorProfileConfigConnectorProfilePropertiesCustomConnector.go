// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_connector_profile#o_auth_2_properties AppflowConnectorProfile#o_auth_2_properties}.
	OAuth2Properties *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnectorOAuth2Properties `field:"optional" json:"oAuth2Properties" yaml:"oAuth2Properties"`
	// A map for properties for custom connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appflow_connector_profile#profile_properties AppflowConnectorProfile#profile_properties}
	ProfileProperties *map[string]*string `field:"optional" json:"profileProperties" yaml:"profileProperties"`
}

