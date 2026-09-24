// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package billingconductorpricingrule


type BillingconductorPricingRuleTiering struct {
	// The set of custom volume tiers for a SKU-scoped TIERING pricing rule.
	//
	// Tiers must start at 0, be contiguous, and the last tier must have no end range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/billingconductor_pricing_rule#custom_tiers BillingconductorPricingRule#custom_tiers}
	CustomTiers interface{} `field:"optional" json:"customTiers" yaml:"customTiers"`
	// The possible customizable free tier configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/billingconductor_pricing_rule#free_tier BillingconductorPricingRule#free_tier}
	FreeTier *BillingconductorPricingRuleTieringFreeTier `field:"optional" json:"freeTier" yaml:"freeTier"`
}

