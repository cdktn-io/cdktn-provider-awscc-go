// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupInstanceLifecyclePolicy struct {
	// Specifies the conditions that trigger instance retention behavior.
	//
	// These triggers determine when instances should move to a ``Retained`` state instead of automatic termination. This allows you to maintain control over instance management when lifecycles transition and operations fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/autoscaling_auto_scaling_group#retention_triggers AutoscalingAutoScalingGroup#retention_triggers}
	RetentionTriggers *AutoscalingAutoScalingGroupInstanceLifecyclePolicyRetentionTriggers `field:"optional" json:"retentionTriggers" yaml:"retentionTriggers"`
}

