// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ec2fleet


type Ec2Ec2FleetReservedCapacityOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_ec2_fleet#allocation_strategy Ec2Ec2Fleet#allocation_strategy}.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_ec2_fleet#capacity_reservation_target Ec2Ec2Fleet#capacity_reservation_target}.
	CapacityReservationTarget *Ec2Ec2FleetReservedCapacityOptionsCapacityReservationTarget `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_ec2_fleet#reservation_types Ec2Ec2Fleet#reservation_types}.
	ReservationTypes *[]*string `field:"optional" json:"reservationTypes" yaml:"reservationTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_ec2_fleet#reserved_capacity_fallback_options Ec2Ec2Fleet#reserved_capacity_fallback_options}.
	ReservedCapacityFallbackOptions *Ec2Ec2FleetReservedCapacityOptionsReservedCapacityFallbackOptions `field:"optional" json:"reservedCapacityFallbackOptions" yaml:"reservedCapacityFallbackOptions"`
}

