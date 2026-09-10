// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupInstanceMaintenancePolicy struct {
	// Specifies the upper threshold as a percentage of the desired capacity of the Auto Scaling group.
	//
	// It represents the maximum percentage of the group that can be in service and healthy, or pending, to support your workload when replacing instances. Value range is 100 to 200. To clear a previously set value, specify a value of ``-1``.
	//  Both ``MinHealthyPercentage`` and ``MaxHealthyPercentage`` must be specified, and the difference between them cannot be greater than 100. A large range increases the number of instances that can be replaced at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_auto_scaling_group#max_healthy_percentage AutoscalingAutoScalingGroup#max_healthy_percentage}
	MaxHealthyPercentage *float64 `field:"optional" json:"maxHealthyPercentage" yaml:"maxHealthyPercentage"`
	// Specifies the lower threshold as a percentage of the desired capacity of the Auto Scaling group.
	//
	// It represents the minimum percentage of the group to keep in service, healthy, and ready to use to support your workload when replacing instances. Value range is 0 to 100. To clear a previously set value, specify a value of ``-1``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_auto_scaling_group#min_healthy_percentage AutoscalingAutoScalingGroup#min_healthy_percentage}
	MinHealthyPercentage *float64 `field:"optional" json:"minHealthyPercentage" yaml:"minHealthyPercentage"`
}

