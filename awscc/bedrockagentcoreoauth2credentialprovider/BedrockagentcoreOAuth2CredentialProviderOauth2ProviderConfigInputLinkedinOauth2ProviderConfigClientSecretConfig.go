// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputLinkedinOauth2ProviderConfigClientSecretConfig struct {
	// The JSON key within the secret that contains the credential value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#json_key BedrockagentcoreOAuth2CredentialProvider#json_key}
	JsonKey *string `field:"optional" json:"jsonKey" yaml:"jsonKey"`
	// The ID or ARN of the secret in AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#secret_id BedrockagentcoreOAuth2CredentialProvider#secret_id}
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
}

