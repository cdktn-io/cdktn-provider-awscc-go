// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsAcceleratorCount struct {
	// The maximum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_auto_scaling_group#max AutoscalingAutoScalingGroup#max}
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_auto_scaling_group#min AutoscalingAutoScalingGroup#min}
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

