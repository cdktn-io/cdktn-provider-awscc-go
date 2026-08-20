// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorepaymentcredentialprovider


type BedrockagentcorePaymentCredentialProviderProviderConfigurationInput struct {
	// Coinbase CDP configuration with API credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_payment_credential_provider#coinbase_cdp_configuration BedrockagentcorePaymentCredentialProvider#coinbase_cdp_configuration}
	CoinbaseCdpConfiguration *BedrockagentcorePaymentCredentialProviderProviderConfigurationInputCoinbaseCdpConfiguration `field:"optional" json:"coinbaseCdpConfiguration" yaml:"coinbaseCdpConfiguration"`
	// Stripe Privy configuration with credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_payment_credential_provider#stripe_privy_configuration BedrockagentcorePaymentCredentialProvider#stripe_privy_configuration}
	StripePrivyConfiguration *BedrockagentcorePaymentCredentialProviderProviderConfigurationInputStripePrivyConfiguration `field:"optional" json:"stripePrivyConfiguration" yaml:"stripePrivyConfiguration"`
}

