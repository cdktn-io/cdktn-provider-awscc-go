// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenacapacityreservation


type AthenaCapacityReservationTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_capacity_reservation#key AthenaCapacityReservation#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_capacity_reservation#value AthenaCapacityReservation#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

