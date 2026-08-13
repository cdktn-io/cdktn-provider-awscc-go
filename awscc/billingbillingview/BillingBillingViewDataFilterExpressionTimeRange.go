// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package billingbillingview


type BillingBillingViewDataFilterExpressionTimeRange struct {
	// The time in ISO 8601 format, UTC time (YYYY-MM-DDTHH:MM:SSZ).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/billing_billing_view#begin_date_inclusive BillingBillingView#begin_date_inclusive}
	BeginDateInclusive *string `field:"optional" json:"beginDateInclusive" yaml:"beginDateInclusive"`
	// The time in ISO 8601 format, UTC time (YYYY-MM-DDTHH:MM:SSZ).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/billing_billing_view#end_date_inclusive BillingBillingView#end_date_inclusive}
	EndDateInclusive *string `field:"optional" json:"endDateInclusive" yaml:"endDateInclusive"`
}

