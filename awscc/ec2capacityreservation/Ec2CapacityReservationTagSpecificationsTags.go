// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2capacityreservation


type Ec2CapacityReservationTagSpecificationsTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_capacity_reservation#key Ec2CapacityReservation#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_capacity_reservation#value Ec2CapacityReservation#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

