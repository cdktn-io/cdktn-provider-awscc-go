// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenacapacityreservation


type AthenaCapacityReservationCapacityAssignmentConfiguration struct {
	// List of capacity assignments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/athena_capacity_reservation#capacity_assignments AthenaCapacityReservation#capacity_assignments}
	CapacityAssignments interface{} `field:"optional" json:"capacityAssignments" yaml:"capacityAssignments"`
}

