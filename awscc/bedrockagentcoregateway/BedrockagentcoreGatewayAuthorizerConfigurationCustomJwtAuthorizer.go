// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregateway


type BedrockagentcoreGatewayAuthorizerConfigurationCustomJwtAuthorizer struct {
	// Maps an originalScope (from allowedScopes) to an advertisedScope exposed in WWW-Authenticate / Protected Resource Metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_gateway#advertised_scope_mapping BedrockagentcoreGateway#advertised_scope_mapping}
	AdvertisedScopeMapping *map[string]*string `field:"optional" json:"advertisedScopeMapping" yaml:"advertisedScopeMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_gateway#allowed_audience BedrockagentcoreGateway#allowed_audience}.
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_gateway#allowed_clients BedrockagentcoreGateway#allowed_clients}.
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_gateway#allowed_scopes BedrockagentcoreGateway#allowed_scopes}.
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_gateway#custom_claims BedrockagentcoreGateway#custom_claims}.
	CustomClaims interface{} `field:"optional" json:"customClaims" yaml:"customClaims"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_gateway#discovery_url BedrockagentcoreGateway#discovery_url}.
	DiscoveryUrl *string `field:"optional" json:"discoveryUrl" yaml:"discoveryUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_gateway#private_endpoint BedrockagentcoreGateway#private_endpoint}.
	PrivateEndpoint *BedrockagentcoreGatewayAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpoint `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
}

