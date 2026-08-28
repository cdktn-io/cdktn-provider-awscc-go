// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorepaymentconnector


type BedrockagentcorePaymentConnectorCredentialProviderConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_payment_connector#coinbase_cdp BedrockagentcorePaymentConnector#coinbase_cdp}.
	CoinbaseCdp *BedrockagentcorePaymentConnectorCredentialProviderConfigurationsCoinbaseCdp `field:"optional" json:"coinbaseCdp" yaml:"coinbaseCdp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_payment_connector#stripe_privy BedrockagentcorePaymentConnector#stripe_privy}.
	StripePrivy *BedrockagentcorePaymentConnectorCredentialProviderConfigurationsStripePrivy `field:"optional" json:"stripePrivy" yaml:"stripePrivy"`
}

