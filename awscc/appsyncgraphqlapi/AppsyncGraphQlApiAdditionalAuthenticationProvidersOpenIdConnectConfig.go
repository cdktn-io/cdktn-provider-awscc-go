// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncgraphqlapi


type AppsyncGraphQlApiAdditionalAuthenticationProvidersOpenIdConnectConfig struct {
	// The number of milliseconds that a token is valid after being authenticated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appsync_graph_ql_api#auth_ttl AppsyncGraphQlApi#auth_ttl}
	AuthTtl *float64 `field:"optional" json:"authTtl" yaml:"authTtl"`
	// The client identifier of the Relying party at the OpenID identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appsync_graph_ql_api#client_id AppsyncGraphQlApi#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The number of milliseconds that a token is valid after it's issued to a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appsync_graph_ql_api#iat_ttl AppsyncGraphQlApi#iat_ttl}
	IatTtl *float64 `field:"optional" json:"iatTtl" yaml:"iatTtl"`
	// The issuer for the OIDC configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appsync_graph_ql_api#issuer AppsyncGraphQlApi#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
}

