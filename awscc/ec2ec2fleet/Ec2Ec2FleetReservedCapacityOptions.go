// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ec2fleet


type Ec2Ec2FleetReservedCapacityOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ec2_fleet#reservation_types Ec2Ec2Fleet#reservation_types}.
	ReservationTypes *[]*string `field:"optional" json:"reservationTypes" yaml:"reservationTypes"`
}

