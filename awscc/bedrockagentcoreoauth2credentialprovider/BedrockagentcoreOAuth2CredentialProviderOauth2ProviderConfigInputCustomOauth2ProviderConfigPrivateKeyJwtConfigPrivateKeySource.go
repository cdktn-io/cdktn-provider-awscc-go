// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigPrivateKeyJwtConfigPrivateKeySource struct {
	// Contains the KMS key configuration for a JWT client assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#kms_key_source BedrockagentcoreOAuth2CredentialProvider#kms_key_source}
	KmsKeySource *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigPrivateKeyJwtConfigPrivateKeySourceKmsKeySource `field:"optional" json:"kmsKeySource" yaml:"kmsKeySource"`
}

