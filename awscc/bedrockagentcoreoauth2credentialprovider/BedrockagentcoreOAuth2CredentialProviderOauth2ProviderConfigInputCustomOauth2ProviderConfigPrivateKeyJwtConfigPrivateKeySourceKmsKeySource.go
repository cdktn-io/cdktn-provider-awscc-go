// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigPrivateKeyJwtConfigPrivateKeySourceKmsKeySource struct {
	// The Amazon Resource Name (ARN) of the KMS key used to sign the JWT client assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#kms_key_arn BedrockagentcoreOAuth2CredentialProvider#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

