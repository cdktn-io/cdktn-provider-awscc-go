// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetsbudgetsaction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BudgetsBudgetsActionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#action_threshold BudgetsBudgetsAction#action_threshold}.
	ActionThreshold *BudgetsBudgetsActionActionThreshold `field:"required" json:"actionThreshold" yaml:"actionThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#action_type BudgetsBudgetsAction#action_type}.
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#budget_name BudgetsBudgetsAction#budget_name}.
	BudgetName *string `field:"required" json:"budgetName" yaml:"budgetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#definition BudgetsBudgetsAction#definition}.
	Definition *BudgetsBudgetsActionDefinition `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#execution_role_arn BudgetsBudgetsAction#execution_role_arn}.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#notification_type BudgetsBudgetsAction#notification_type}.
	NotificationType *string `field:"required" json:"notificationType" yaml:"notificationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#subscribers BudgetsBudgetsAction#subscribers}.
	Subscribers interface{} `field:"required" json:"subscribers" yaml:"subscribers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#approval_model BudgetsBudgetsAction#approval_model}.
	ApprovalModel *string `field:"optional" json:"approvalModel" yaml:"approvalModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/budgets_budgets_action#resource_tags BudgetsBudgetsAction#resource_tags}.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

