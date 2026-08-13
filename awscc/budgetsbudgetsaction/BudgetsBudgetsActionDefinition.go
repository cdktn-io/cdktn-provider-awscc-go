// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetsbudgetsaction


type BudgetsBudgetsActionDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/budgets_budgets_action#iam_action_definition BudgetsBudgetsAction#iam_action_definition}.
	IamActionDefinition *BudgetsBudgetsActionDefinitionIamActionDefinition `field:"optional" json:"iamActionDefinition" yaml:"iamActionDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/budgets_budgets_action#scp_action_definition BudgetsBudgetsAction#scp_action_definition}.
	ScpActionDefinition *BudgetsBudgetsActionDefinitionScpActionDefinition `field:"optional" json:"scpActionDefinition" yaml:"scpActionDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/budgets_budgets_action#ssm_action_definition BudgetsBudgetsAction#ssm_action_definition}.
	SsmActionDefinition *BudgetsBudgetsActionDefinitionSsmActionDefinition `field:"optional" json:"ssmActionDefinition" yaml:"ssmActionDefinition"`
}

