// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigOnBehalfOfTokenExchangeConfig struct {
	// The grant type for on-behalf-of token exchange.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#grant_type BedrockagentcoreOAuth2CredentialProvider#grant_type}
	GrantType *string `field:"optional" json:"grantType" yaml:"grantType"`
	// Configuration for RFC 8693 Token Exchange.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#token_exchange_grant_type_config BedrockagentcoreOAuth2CredentialProvider#token_exchange_grant_type_config}
	TokenExchangeGrantTypeConfig *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigOnBehalfOfTokenExchangeConfigTokenExchangeGrantTypeConfig `field:"optional" json:"tokenExchangeGrantTypeConfig" yaml:"tokenExchangeGrantTypeConfig"`
}

