// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigOAuthClientCredentials struct {
	// OAuth client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_service#client_id DevopsagentService#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// User friendly OAuth client name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_service#client_name DevopsagentService#client_name}
	ClientName *string `field:"optional" json:"clientName" yaml:"clientName"`
	// OAuth client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_service#client_secret DevopsagentService#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// OAuth token exchange parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_service#exchange_parameters DevopsagentService#exchange_parameters}
	ExchangeParameters *string `field:"optional" json:"exchangeParameters" yaml:"exchangeParameters"`
	// OAuth token exchange URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_service#exchange_url DevopsagentService#exchange_url}
	ExchangeUrl *string `field:"optional" json:"exchangeUrl" yaml:"exchangeUrl"`
	// OAuth scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_service#scopes DevopsagentService#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

