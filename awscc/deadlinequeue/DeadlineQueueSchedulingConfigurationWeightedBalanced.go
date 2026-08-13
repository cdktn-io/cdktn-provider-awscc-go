// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinequeue


type DeadlineQueueSchedulingConfigurationWeightedBalanced struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/deadline_queue#error_weight DeadlineQueue#error_weight}.
	ErrorWeight *float64 `field:"optional" json:"errorWeight" yaml:"errorWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/deadline_queue#max_priority_override DeadlineQueue#max_priority_override}.
	MaxPriorityOverride *DeadlineQueueSchedulingConfigurationWeightedBalancedMaxPriorityOverride `field:"optional" json:"maxPriorityOverride" yaml:"maxPriorityOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/deadline_queue#min_priority_override DeadlineQueue#min_priority_override}.
	MinPriorityOverride *DeadlineQueueSchedulingConfigurationWeightedBalancedMinPriorityOverride `field:"optional" json:"minPriorityOverride" yaml:"minPriorityOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/deadline_queue#priority_weight DeadlineQueue#priority_weight}.
	PriorityWeight *float64 `field:"optional" json:"priorityWeight" yaml:"priorityWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/deadline_queue#rendering_task_buffer DeadlineQueue#rendering_task_buffer}.
	RenderingTaskBuffer *float64 `field:"optional" json:"renderingTaskBuffer" yaml:"renderingTaskBuffer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/deadline_queue#rendering_task_weight DeadlineQueue#rendering_task_weight}.
	RenderingTaskWeight *float64 `field:"optional" json:"renderingTaskWeight" yaml:"renderingTaskWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/deadline_queue#submission_time_weight DeadlineQueue#submission_time_weight}.
	SubmissionTimeWeight *float64 `field:"optional" json:"submissionTimeWeight" yaml:"submissionTimeWeight"`
}

