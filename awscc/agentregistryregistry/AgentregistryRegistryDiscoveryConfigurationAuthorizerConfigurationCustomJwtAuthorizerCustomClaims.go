// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistry


type AgentregistryRegistryDiscoveryConfigurationAuthorizerConfigurationCustomJwtAuthorizerCustomClaims struct {
	// The value and match operator used to authorize a claim during JWT validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/agentregistry_registry#authorizing_claim_match_value AgentregistryRegistry#authorizing_claim_match_value}
	AuthorizingClaimMatchValue *AgentregistryRegistryDiscoveryConfigurationAuthorizerConfigurationCustomJwtAuthorizerCustomClaimsAuthorizingClaimMatchValue `field:"optional" json:"authorizingClaimMatchValue" yaml:"authorizingClaimMatchValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/agentregistry_registry#inbound_token_claim_name AgentregistryRegistry#inbound_token_claim_name}.
	InboundTokenClaimName *string `field:"optional" json:"inboundTokenClaimName" yaml:"inboundTokenClaimName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/agentregistry_registry#inbound_token_claim_value_type AgentregistryRegistry#inbound_token_claim_value_type}.
	InboundTokenClaimValueType *string `field:"optional" json:"inboundTokenClaimValueType" yaml:"inboundTokenClaimValueType"`
}

