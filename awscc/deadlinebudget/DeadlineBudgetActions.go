// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinebudget


type DeadlineBudgetActions struct {
	// The percentage threshold for the budget action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_budget#threshold_percentage DeadlineBudget#threshold_percentage}
	ThresholdPercentage *float64 `field:"required" json:"thresholdPercentage" yaml:"thresholdPercentage"`
	// The type of budget action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_budget#type DeadlineBudget#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description for the budget action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_budget#description DeadlineBudget#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

