// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package billingconductorcustomlineitem


type BillingconductorCustomLineItemBillingPeriodRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/billingconductor_custom_line_item#exclusive_end_billing_period BillingconductorCustomLineItem#exclusive_end_billing_period}.
	ExclusiveEndBillingPeriod *string `field:"optional" json:"exclusiveEndBillingPeriod" yaml:"exclusiveEndBillingPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/billingconductor_custom_line_item#inclusive_start_billing_period BillingconductorCustomLineItem#inclusive_start_billing_period}.
	InclusiveStartBillingPeriod *string `field:"optional" json:"inclusiveStartBillingPeriod" yaml:"inclusiveStartBillingPeriod"`
}

