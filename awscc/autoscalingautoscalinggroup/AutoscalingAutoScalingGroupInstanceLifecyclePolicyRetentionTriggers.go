// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupInstanceLifecyclePolicyRetentionTriggers struct {
	// Specifies the action when a termination lifecycle hook is abandoned due to failure, timeout, or explicit abandonment (calling CompleteLifecycleAction).
	//
	// Set to ``retain`` to move instances to a retained state. Set to ``terminate`` for default termination behavior.
	//   Retained instances don't count toward desired capacity and remain until you call ``TerminateInstanceInAutoScalingGroup``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_auto_scaling_group#terminate_hook_abandon AutoscalingAutoScalingGroup#terminate_hook_abandon}
	TerminateHookAbandon *string `field:"optional" json:"terminateHookAbandon" yaml:"terminateHookAbandon"`
}

