// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenacapacityreservation


type AthenaCapacityReservationCapacityAssignmentConfigurationCapacityAssignments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_capacity_reservation#workgroup_names AthenaCapacityReservation#workgroup_names}.
	WorkgroupNames *[]*string `field:"optional" json:"workgroupNames" yaml:"workgroupNames"`
}

