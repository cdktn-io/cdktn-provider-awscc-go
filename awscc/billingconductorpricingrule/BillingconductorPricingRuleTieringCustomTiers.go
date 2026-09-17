// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package billingconductorpricingrule


type BillingconductorPricingRuleTieringCustomTiers struct {
	// The inclusive beginning of the tier's usage range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/billingconductor_pricing_rule#begin_range_inclusive BillingconductorPricingRule#begin_range_inclusive}
	BeginRangeInclusive *float64 `field:"optional" json:"beginRangeInclusive" yaml:"beginRangeInclusive"`
	// The exclusive end of the tier's usage range. Omit for the last tier (infinity).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/billingconductor_pricing_rule#end_range_exclusive BillingconductorPricingRule#end_range_exclusive}
	EndRangeExclusive *float64 `field:"optional" json:"endRangeExclusive" yaml:"endRangeExclusive"`
	// The custom rate applied to usage within the tier's range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/billingconductor_pricing_rule#rate_value BillingconductorPricingRule#rate_value}
	RateValue *float64 `field:"optional" json:"rateValue" yaml:"rateValue"`
}

