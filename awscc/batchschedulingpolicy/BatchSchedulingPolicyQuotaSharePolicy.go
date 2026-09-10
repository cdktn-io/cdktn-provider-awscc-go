// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchschedulingpolicy


type BatchSchedulingPolicyQuotaSharePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_scheduling_policy#idle_resource_assignment_strategy BatchSchedulingPolicy#idle_resource_assignment_strategy}.
	IdleResourceAssignmentStrategy *string `field:"optional" json:"idleResourceAssignmentStrategy" yaml:"idleResourceAssignmentStrategy"`
}

