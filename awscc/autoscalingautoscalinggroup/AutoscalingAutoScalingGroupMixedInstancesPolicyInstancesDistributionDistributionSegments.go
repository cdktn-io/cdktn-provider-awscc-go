// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionDistributionSegments struct {
	// The capacity types to prioritize, in order.
	//
	// Amazon EC2 Auto Scaling attempts to launch instances in the priority order of the capacity types, and within each capacity type, in the order of instance types listed in your launch template ``Overrides``.
	//  The following lists the valid values:
	//   + on-demand-capacity-reservation On-Demand Capacity Reservations. + capacity-block Capacity Blocks. + interruptible-capacity-reservation Interruptible Capacity Reservations. + on-demand On-Demand capacity. Include this value to allow the group to fall back to On-Demand capacity when the preceding capacity types are unavailable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_auto_scaling_group#target_capacity_types AutoscalingAutoScalingGroup#target_capacity_types}
	TargetCapacityTypes *[]*string `field:"optional" json:"targetCapacityTypes" yaml:"targetCapacityTypes"`
}

