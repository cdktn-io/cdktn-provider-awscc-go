// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetsbudgetsaction


type BudgetsBudgetsActionDefinitionSsmActionDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#instance_ids BudgetsBudgetsAction#instance_ids}.
	InstanceIds *[]*string `field:"optional" json:"instanceIds" yaml:"instanceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#region BudgetsBudgetsAction#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#subtype BudgetsBudgetsAction#subtype}.
	Subtype *string `field:"optional" json:"subtype" yaml:"subtype"`
}

