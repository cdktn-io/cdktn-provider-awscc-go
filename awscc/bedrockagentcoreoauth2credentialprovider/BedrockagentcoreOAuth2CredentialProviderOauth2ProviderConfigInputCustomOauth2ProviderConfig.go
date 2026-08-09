// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfig struct {
	// The client authentication method to use when authenticating with the token endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#client_authentication_method BedrockagentcoreOAuth2CredentialProvider#client_authentication_method}
	ClientAuthenticationMethod *string `field:"optional" json:"clientAuthenticationMethod" yaml:"clientAuthenticationMethod"`
	// The client ID for the custom OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#client_id BedrockagentcoreOAuth2CredentialProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client secret for the custom OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#client_secret BedrockagentcoreOAuth2CredentialProvider#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// A reference to a customer-provided secret stored in AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#client_secret_config BedrockagentcoreOAuth2CredentialProvider#client_secret_config}
	ClientSecretConfig *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigClientSecretConfig `field:"optional" json:"clientSecretConfig" yaml:"clientSecretConfig"`
	// The source of the client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#client_secret_source BedrockagentcoreOAuth2CredentialProvider#client_secret_source}
	ClientSecretSource *string `field:"optional" json:"clientSecretSource" yaml:"clientSecretSource"`
	// Discovery information for an OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#oauth_discovery BedrockagentcoreOAuth2CredentialProvider#oauth_discovery}
	OauthDiscovery *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigOauthDiscovery `field:"optional" json:"oauthDiscovery" yaml:"oauthDiscovery"`
	// Configuration for on-behalf-of token exchange.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#on_behalf_of_token_exchange_config BedrockagentcoreOAuth2CredentialProvider#on_behalf_of_token_exchange_config}
	OnBehalfOfTokenExchangeConfig *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigOnBehalfOfTokenExchangeConfig `field:"optional" json:"onBehalfOfTokenExchangeConfig" yaml:"onBehalfOfTokenExchangeConfig"`
	// The private endpoint configuration for connecting to private resources in your VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#private_endpoint BedrockagentcoreOAuth2CredentialProvider#private_endpoint}
	PrivateEndpoint *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigPrivateEndpoint `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
	// A list of private endpoint overrides.
	//
	// Each override maps a specific domain to a private endpoint, enabling secure connectivity through VPC Lattice resource configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#private_endpoint_overrides BedrockagentcoreOAuth2CredentialProvider#private_endpoint_overrides}
	PrivateEndpointOverrides interface{} `field:"optional" json:"privateEndpointOverrides" yaml:"privateEndpointOverrides"`
	// Configuration for private_key_jwt client authentication (RFC 7523).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#private_key_jwt_config BedrockagentcoreOAuth2CredentialProvider#private_key_jwt_config}
	PrivateKeyJwtConfig *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigPrivateKeyJwtConfig `field:"optional" json:"privateKeyJwtConfig" yaml:"privateKeyJwtConfig"`
}

