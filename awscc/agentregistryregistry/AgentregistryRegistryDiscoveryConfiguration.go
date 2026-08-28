// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistry


type AgentregistryRegistryDiscoveryConfiguration struct {
	// The authorizer configuration for the registry. This is a union - specify exactly one member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/agentregistry_registry#authorizer_configuration AgentregistryRegistry#authorizer_configuration}
	AuthorizerConfiguration *AgentregistryRegistryDiscoveryConfigurationAuthorizerConfiguration `field:"optional" json:"authorizerConfiguration" yaml:"authorizerConfiguration"`
}

