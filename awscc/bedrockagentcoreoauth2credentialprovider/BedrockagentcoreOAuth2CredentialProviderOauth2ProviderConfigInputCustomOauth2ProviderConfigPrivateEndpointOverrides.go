// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigPrivateEndpointOverrides struct {
	// The domain to override with a private endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#domain BedrockagentcoreOAuth2CredentialProvider#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The private endpoint configuration for connecting to private resources in your VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#private_endpoint BedrockagentcoreOAuth2CredentialProvider#private_endpoint}
	PrivateEndpoint *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigPrivateEndpointOverridesPrivateEndpoint `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
}

