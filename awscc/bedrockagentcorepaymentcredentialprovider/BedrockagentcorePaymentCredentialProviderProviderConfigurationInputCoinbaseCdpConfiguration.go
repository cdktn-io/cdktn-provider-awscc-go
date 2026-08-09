// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorepaymentcredentialprovider


type BedrockagentcorePaymentCredentialProviderProviderConfigurationInputCoinbaseCdpConfiguration struct {
	// The Coinbase CDP API key ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_payment_credential_provider#api_key_id BedrockagentcorePaymentCredentialProvider#api_key_id}
	ApiKeyId *string `field:"optional" json:"apiKeyId" yaml:"apiKeyId"`
	// The Coinbase CDP API key secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_payment_credential_provider#api_key_secret BedrockagentcorePaymentCredentialProvider#api_key_secret}
	ApiKeySecret *string `field:"optional" json:"apiKeySecret" yaml:"apiKeySecret"`
	// A reference to a customer-provided secret stored in AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_payment_credential_provider#api_key_secret_config BedrockagentcorePaymentCredentialProvider#api_key_secret_config}
	ApiKeySecretConfig *BedrockagentcorePaymentCredentialProviderProviderConfigurationInputCoinbaseCdpConfigurationApiKeySecretConfig `field:"optional" json:"apiKeySecretConfig" yaml:"apiKeySecretConfig"`
	// The source of the secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_payment_credential_provider#api_key_secret_source BedrockagentcorePaymentCredentialProvider#api_key_secret_source}
	ApiKeySecretSource *string `field:"optional" json:"apiKeySecretSource" yaml:"apiKeySecretSource"`
	// The Coinbase CDP wallet secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_payment_credential_provider#wallet_secret BedrockagentcorePaymentCredentialProvider#wallet_secret}
	WalletSecret *string `field:"optional" json:"walletSecret" yaml:"walletSecret"`
	// A reference to a customer-provided secret stored in AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_payment_credential_provider#wallet_secret_config BedrockagentcorePaymentCredentialProvider#wallet_secret_config}
	WalletSecretConfig *BedrockagentcorePaymentCredentialProviderProviderConfigurationInputCoinbaseCdpConfigurationWalletSecretConfig `field:"optional" json:"walletSecretConfig" yaml:"walletSecretConfig"`
	// The source of the secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_payment_credential_provider#wallet_secret_source BedrockagentcorePaymentCredentialProvider#wallet_secret_source}
	WalletSecretSource *string `field:"optional" json:"walletSecretSource" yaml:"walletSecretSource"`
}

