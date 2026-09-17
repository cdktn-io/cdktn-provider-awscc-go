// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_connector_profile#application_host_url AppflowConnectorProfile#application_host_url}.
	ApplicationHostUrl *string `field:"optional" json:"applicationHostUrl" yaml:"applicationHostUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_connector_profile#application_service_path AppflowConnectorProfile#application_service_path}.
	ApplicationServicePath *string `field:"optional" json:"applicationServicePath" yaml:"applicationServicePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_connector_profile#client_number AppflowConnectorProfile#client_number}.
	ClientNumber *string `field:"optional" json:"clientNumber" yaml:"clientNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_connector_profile#logon_language AppflowConnectorProfile#logon_language}.
	LogonLanguage *string `field:"optional" json:"logonLanguage" yaml:"logonLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_connector_profile#o_auth_properties AppflowConnectorProfile#o_auth_properties}.
	OAuthProperties *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOAuthProperties `field:"optional" json:"oAuthProperties" yaml:"oAuthProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_connector_profile#port_number AppflowConnectorProfile#port_number}.
	PortNumber *float64 `field:"optional" json:"portNumber" yaml:"portNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_connector_profile#private_link_service_name AppflowConnectorProfile#private_link_service_name}.
	PrivateLinkServiceName *string `field:"optional" json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
}

