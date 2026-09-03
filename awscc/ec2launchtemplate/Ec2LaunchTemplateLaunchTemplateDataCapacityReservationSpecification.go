// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecification struct {
	// Indicates the instance's Capacity Reservation preferences.
	//
	// Possible preferences include:
	//   +   ``open`` - The instance can run in any ``open`` Capacity Reservation that has matching attributes (instance type, platform, Availability Zone).
	//   +   ``none`` - The instance avoids running in a Capacity Reservation even if one is available. The instance runs in On-Demand capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_launch_template#capacity_reservation_preference Ec2LaunchTemplate#capacity_reservation_preference}
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// Information about the target Capacity Reservation or Capacity Reservation group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_launch_template#capacity_reservation_target Ec2LaunchTemplate#capacity_reservation_target}
	CapacityReservationTarget *Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationCapacityReservationTarget `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

