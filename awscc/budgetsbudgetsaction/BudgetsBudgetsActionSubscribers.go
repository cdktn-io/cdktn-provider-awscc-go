// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetsbudgetsaction


type BudgetsBudgetsActionSubscribers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#address BudgetsBudgetsAction#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#type BudgetsBudgetsAction#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

