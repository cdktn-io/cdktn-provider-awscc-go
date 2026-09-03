// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigOauthDiscovery struct {
	// Authorization server metadata for the OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#authorization_server_metadata BedrockagentcoreOAuth2CredentialProvider#authorization_server_metadata}
	AuthorizationServerMetadata *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigOauthDiscoveryAuthorizationServerMetadata `field:"optional" json:"authorizationServerMetadata" yaml:"authorizationServerMetadata"`
	// The discovery URL for the OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#discovery_url BedrockagentcoreOAuth2CredentialProvider#discovery_url}
	DiscoveryUrl *string `field:"optional" json:"discoveryUrl" yaml:"discoveryUrl"`
}

