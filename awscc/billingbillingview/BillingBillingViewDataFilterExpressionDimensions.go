// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package billingbillingview


type BillingBillingViewDataFilterExpressionDimensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/billing_billing_view#key BillingBillingView#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/billing_billing_view#values BillingBillingView#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

