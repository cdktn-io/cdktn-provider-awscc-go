// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsBaselinePerformanceFactorsCpu struct {
	// Specify an instance family to use as the baseline reference for CPU performance.
	//
	// All instance types that match your specified attributes will be compared against the CPU performance of the referenced instance family, regardless of CPU manufacturer or architecture differences.
	//   Currently only one instance family can be specified in the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/autoscaling_auto_scaling_group#references AutoscalingAutoScalingGroup#references}
	References interface{} `field:"optional" json:"references" yaml:"references"`
}

