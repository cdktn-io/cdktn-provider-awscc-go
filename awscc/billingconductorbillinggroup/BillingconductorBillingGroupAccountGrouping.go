// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package billingconductorbillinggroup


type BillingconductorBillingGroupAccountGrouping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/billingconductor_billing_group#auto_associate BillingconductorBillingGroup#auto_associate}.
	AutoAssociate interface{} `field:"optional" json:"autoAssociate" yaml:"autoAssociate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/billingconductor_billing_group#linked_account_ids BillingconductorBillingGroup#linked_account_ids}.
	LinkedAccountIds *[]*string `field:"optional" json:"linkedAccountIds" yaml:"linkedAccountIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/billingconductor_billing_group#responsibility_transfer_arn BillingconductorBillingGroup#responsibility_transfer_arn}.
	ResponsibilityTransferArn *string `field:"optional" json:"responsibilityTransferArn" yaml:"responsibilityTransferArn"`
}

