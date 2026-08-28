// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistry


type AgentregistryRegistryDiscoveryConfigurationAuthorizerConfigurationCustomJwtAuthorizer struct {
	// The audience values accepted during JWT validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/agentregistry_registry#allowed_audience AgentregistryRegistry#allowed_audience}
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// The client identifiers accepted during JWT validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/agentregistry_registry#allowed_clients AgentregistryRegistry#allowed_clients}
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// The scopes accepted during JWT validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/agentregistry_registry#allowed_scopes AgentregistryRegistry#allowed_scopes}
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// Additional custom claim validations applied to the inbound JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/agentregistry_registry#custom_claims AgentregistryRegistry#custom_claims}
	CustomClaims interface{} `field:"optional" json:"customClaims" yaml:"customClaims"`
	// The OpenID Connect discovery URL used to retrieve the identity provider's metadata and signing keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/agentregistry_registry#discovery_url AgentregistryRegistry#discovery_url}
	DiscoveryUrl *string `field:"optional" json:"discoveryUrl" yaml:"discoveryUrl"`
}

