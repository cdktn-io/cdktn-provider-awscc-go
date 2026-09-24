// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinequeue


type DeadlineQueueSchedulingConfigurationWeightedBalancedMinPriorityOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/deadline_queue#always_schedule_last DeadlineQueue#always_schedule_last}.
	AlwaysScheduleLast *string `field:"optional" json:"alwaysScheduleLast" yaml:"alwaysScheduleLast"`
}

