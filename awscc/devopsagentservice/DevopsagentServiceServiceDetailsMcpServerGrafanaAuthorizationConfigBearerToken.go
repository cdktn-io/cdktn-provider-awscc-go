// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsMcpServerGrafanaAuthorizationConfigBearerToken struct {
	// HTTP header name to send the bearer token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#authorization_header DevopsagentService#authorization_header}
	AuthorizationHeader *string `field:"optional" json:"authorizationHeader" yaml:"authorizationHeader"`
	// User friendly bearer token name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#token_name DevopsagentService#token_name}
	TokenName *string `field:"optional" json:"tokenName" yaml:"tokenName"`
	// Bearer token value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#token_value DevopsagentService#token_value}
	TokenValue *string `field:"optional" json:"tokenValue" yaml:"tokenValue"`
}

