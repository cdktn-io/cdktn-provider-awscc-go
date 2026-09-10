// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigPrivateKeyJwtConfig struct {
	// A map of additional claims to include in the JWT client assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#additional_header_claims BedrockagentcoreOAuth2CredentialProvider#additional_header_claims}
	AdditionalHeaderClaims *map[string]*string `field:"optional" json:"additionalHeaderClaims" yaml:"additionalHeaderClaims"`
	// A map of additional claims to include in the JWT client assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#additional_payload_claims BedrockagentcoreOAuth2CredentialProvider#additional_payload_claims}
	AdditionalPayloadClaims *map[string]*string `field:"optional" json:"additionalPayloadClaims" yaml:"additionalPayloadClaims"`
	// Contains the private key source configuration for a JWT client assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#private_key_source BedrockagentcoreOAuth2CredentialProvider#private_key_source}
	PrivateKeySource *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigPrivateKeyJwtConfigPrivateKeySource `field:"optional" json:"privateKeySource" yaml:"privateKeySource"`
	// The algorithm used to sign the JWT client assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#signing_algorithm BedrockagentcoreOAuth2CredentialProvider#signing_algorithm}
	SigningAlgorithm *string `field:"optional" json:"signingAlgorithm" yaml:"signingAlgorithm"`
}

