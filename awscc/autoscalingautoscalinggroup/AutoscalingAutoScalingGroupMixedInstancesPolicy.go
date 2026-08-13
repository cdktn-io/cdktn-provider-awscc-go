// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMixedInstancesPolicy struct {
	// The instances distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/autoscaling_auto_scaling_group#instances_distribution AutoscalingAutoScalingGroup#instances_distribution}
	InstancesDistribution *AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistribution `field:"optional" json:"instancesDistribution" yaml:"instancesDistribution"`
	// One or more launch templates and the instance types (overrides) that are used to launch EC2 instances to fulfill On-Demand and Spot capacities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/autoscaling_auto_scaling_group#launch_template AutoscalingAutoScalingGroup#launch_template}
	LaunchTemplate *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
}

