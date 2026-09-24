// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigOnBehalfOfTokenExchangeConfigTokenExchangeGrantTypeConfig struct {
	// The actor token content type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#actor_token_content BedrockagentcoreOAuth2CredentialProvider#actor_token_content}
	ActorTokenContent *string `field:"optional" json:"actorTokenContent" yaml:"actorTokenContent"`
	// The actor token scopes. Only valid when ActorTokenContent is M2M.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#actor_token_scopes BedrockagentcoreOAuth2CredentialProvider#actor_token_scopes}
	ActorTokenScopes *[]*string `field:"optional" json:"actorTokenScopes" yaml:"actorTokenScopes"`
}

