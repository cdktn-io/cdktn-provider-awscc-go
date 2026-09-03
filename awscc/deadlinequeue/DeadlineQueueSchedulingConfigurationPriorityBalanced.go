// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinequeue


type DeadlineQueueSchedulingConfigurationPriorityBalanced struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/deadline_queue#rendering_task_buffer DeadlineQueue#rendering_task_buffer}.
	RenderingTaskBuffer *float64 `field:"optional" json:"renderingTaskBuffer" yaml:"renderingTaskBuffer"`
}

