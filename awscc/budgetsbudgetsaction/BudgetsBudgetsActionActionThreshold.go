// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetsbudgetsaction


type BudgetsBudgetsActionActionThreshold struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/budgets_budgets_action#type BudgetsBudgetsAction#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/budgets_budgets_action#value BudgetsBudgetsAction#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

