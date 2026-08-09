// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinequeue


type DeadlineQueueSchedulingConfigurationWeightedBalancedMaxPriorityOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/deadline_queue#always_schedule_first DeadlineQueue#always_schedule_first}.
	AlwaysScheduleFirst *string `field:"optional" json:"alwaysScheduleFirst" yaml:"alwaysScheduleFirst"`
}

