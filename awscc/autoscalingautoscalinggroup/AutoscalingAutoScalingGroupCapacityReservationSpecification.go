// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupCapacityReservationSpecification struct {
	// The capacity reservation preference.
	//
	// The following options are available:
	//   +  ``capacity-reservations-only`` - Auto Scaling will only launch instances into a Capacity Reservation or Capacity Reservation resource group. If capacity isn't available, instances will fail to launch.
	//   +  ``capacity-reservations-first`` - Auto Scaling will try to launch instances into a Capacity Reservation or Capacity Reservation resource group first. If capacity isn't available, instances will run in On-Demand capacity.
	//   +  ``none`` - Auto Scaling will not launch instances into a Capacity Reservation. Instances will run in On-Demand capacity.
	//   +  ``default`` - Auto Scaling uses the Capacity Reservation preference from your launch template or an open Capacity Reservation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_auto_scaling_group#capacity_reservation_preference AutoscalingAutoScalingGroup#capacity_reservation_preference}
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// Describes a target Capacity Reservation or Capacity Reservation resource group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_auto_scaling_group#capacity_reservation_target AutoscalingAutoScalingGroup#capacity_reservation_target}
	CapacityReservationTarget *AutoscalingAutoScalingGroupCapacityReservationSpecificationCapacityReservationTarget `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

