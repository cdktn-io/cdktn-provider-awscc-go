// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinebudget


type DeadlineBudgetSchedule struct {
	// The details of a fixed budget schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/deadline_budget#fixed DeadlineBudget#fixed}
	Fixed *DeadlineBudgetScheduleFixed `field:"required" json:"fixed" yaml:"fixed"`
}

