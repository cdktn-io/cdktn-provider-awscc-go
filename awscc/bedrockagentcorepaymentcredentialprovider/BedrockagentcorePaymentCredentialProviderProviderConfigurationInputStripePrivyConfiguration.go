// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorepaymentcredentialprovider


type BedrockagentcorePaymentCredentialProviderProviderConfigurationInputStripePrivyConfiguration struct {
	// The app ID provided by Privy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_payment_credential_provider#app_id BedrockagentcorePaymentCredentialProvider#app_id}
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// The app secret provided by Privy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_payment_credential_provider#app_secret BedrockagentcorePaymentCredentialProvider#app_secret}
	AppSecret *string `field:"optional" json:"appSecret" yaml:"appSecret"`
	// A reference to a customer-provided secret stored in AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_payment_credential_provider#app_secret_config BedrockagentcorePaymentCredentialProvider#app_secret_config}
	AppSecretConfig *BedrockagentcorePaymentCredentialProviderProviderConfigurationInputStripePrivyConfigurationAppSecretConfig `field:"optional" json:"appSecretConfig" yaml:"appSecretConfig"`
	// The source of the secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_payment_credential_provider#app_secret_source BedrockagentcorePaymentCredentialProvider#app_secret_source}
	AppSecretSource *string `field:"optional" json:"appSecretSource" yaml:"appSecretSource"`
	// The authorization ID for the Stripe Privy integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_payment_credential_provider#authorization_id BedrockagentcorePaymentCredentialProvider#authorization_id}
	AuthorizationId *string `field:"optional" json:"authorizationId" yaml:"authorizationId"`
	// The authorization private key for the Stripe Privy integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_payment_credential_provider#authorization_private_key BedrockagentcorePaymentCredentialProvider#authorization_private_key}
	AuthorizationPrivateKey *string `field:"optional" json:"authorizationPrivateKey" yaml:"authorizationPrivateKey"`
	// A reference to a customer-provided secret stored in AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_payment_credential_provider#authorization_private_key_config BedrockagentcorePaymentCredentialProvider#authorization_private_key_config}
	AuthorizationPrivateKeyConfig *BedrockagentcorePaymentCredentialProviderProviderConfigurationInputStripePrivyConfigurationAuthorizationPrivateKeyConfig `field:"optional" json:"authorizationPrivateKeyConfig" yaml:"authorizationPrivateKeyConfig"`
	// The source of the secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_payment_credential_provider#authorization_private_key_source BedrockagentcorePaymentCredentialProvider#authorization_private_key_source}
	AuthorizationPrivateKeySource *string `field:"optional" json:"authorizationPrivateKeySource" yaml:"authorizationPrivateKeySource"`
}

