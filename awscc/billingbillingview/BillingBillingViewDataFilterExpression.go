// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package billingbillingview


type BillingBillingViewDataFilterExpression struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/billing_billing_view#dimensions BillingBillingView#dimensions}.
	Dimensions *BillingBillingViewDataFilterExpressionDimensions `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/billing_billing_view#tags BillingBillingView#tags}.
	Tags *BillingBillingViewDataFilterExpressionTags `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/billing_billing_view#time_range BillingBillingView#time_range}.
	TimeRange *BillingBillingViewDataFilterExpressionTimeRange `field:"optional" json:"timeRange" yaml:"timeRange"`
}

