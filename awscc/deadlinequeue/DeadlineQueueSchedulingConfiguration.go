// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinequeue


type DeadlineQueueSchedulingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/deadline_queue#priority_balanced DeadlineQueue#priority_balanced}.
	PriorityBalanced *DeadlineQueueSchedulingConfigurationPriorityBalanced `field:"optional" json:"priorityBalanced" yaml:"priorityBalanced"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/deadline_queue#priority_fifo DeadlineQueue#priority_fifo}.
	PriorityFifo *string `field:"optional" json:"priorityFifo" yaml:"priorityFifo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/deadline_queue#weighted_balanced DeadlineQueue#weighted_balanced}.
	WeightedBalanced *DeadlineQueueSchedulingConfigurationWeightedBalanced `field:"optional" json:"weightedBalanced" yaml:"weightedBalanced"`
}

