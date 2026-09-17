// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment


type BatchComputeEnvironmentComputeResourcesManagedInstancesProviderInstanceLaunchTemplateCapacityReservations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/batch_compute_environment#reservation_group_arn BatchComputeEnvironment#reservation_group_arn}.
	ReservationGroupArn *string `field:"optional" json:"reservationGroupArn" yaml:"reservationGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/batch_compute_environment#reservation_preference BatchComputeEnvironment#reservation_preference}.
	ReservationPreference *string `field:"optional" json:"reservationPreference" yaml:"reservationPreference"`
}

