// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigOauthDiscoveryAuthorizationServerMetadata struct {
	// The authorization endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#authorization_endpoint BedrockagentcoreOAuth2CredentialProvider#authorization_endpoint}
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The issuer URL for the OAuth2 authorization server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#issuer BedrockagentcoreOAuth2CredentialProvider#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// The supported response types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#response_types BedrockagentcoreOAuth2CredentialProvider#response_types}
	ResponseTypes *[]*string `field:"optional" json:"responseTypes" yaml:"responseTypes"`
	// The token endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#token_endpoint BedrockagentcoreOAuth2CredentialProvider#token_endpoint}
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
}

