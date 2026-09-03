// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinebudget


type DeadlineBudgetUsageTrackingResource struct {
	// The queue ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/deadline_budget#queue_id DeadlineBudget#queue_id}
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

