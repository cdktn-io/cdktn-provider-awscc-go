// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupCapacityReservationSpecificationCapacityReservationTarget struct {
	// The Capacity Reservation IDs to launch instances into.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_auto_scaling_group#capacity_reservation_ids AutoscalingAutoScalingGroup#capacity_reservation_ids}
	CapacityReservationIds *[]*string `field:"optional" json:"capacityReservationIds" yaml:"capacityReservationIds"`
	// The resource group ARNs of the Capacity Reservation to launch instances into.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_auto_scaling_group#capacity_reservation_resource_group_arns AutoscalingAutoScalingGroup#capacity_reservation_resource_group_arns}
	CapacityReservationResourceGroupArns *[]*string `field:"optional" json:"capacityReservationResourceGroupArns" yaml:"capacityReservationResourceGroupArns"`
}

