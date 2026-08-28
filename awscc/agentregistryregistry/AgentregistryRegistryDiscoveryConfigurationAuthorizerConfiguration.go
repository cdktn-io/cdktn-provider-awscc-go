// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistry


type AgentregistryRegistryDiscoveryConfigurationAuthorizerConfiguration struct {
	// Configuration for a custom JWT authorizer that validates inbound bearer tokens against an OpenID Connect identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/agentregistry_registry#custom_jwt_authorizer AgentregistryRegistry#custom_jwt_authorizer}
	CustomJwtAuthorizer *AgentregistryRegistryDiscoveryConfigurationAuthorizerConfigurationCustomJwtAuthorizer `field:"optional" json:"customJwtAuthorizer" yaml:"customJwtAuthorizer"`
}

