// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinebudget


type DeadlineBudgetScheduleFixed struct {
	// When the budget ends.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_budget#end_time DeadlineBudget#end_time}
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// When the budget starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_budget#start_time DeadlineBudget#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

