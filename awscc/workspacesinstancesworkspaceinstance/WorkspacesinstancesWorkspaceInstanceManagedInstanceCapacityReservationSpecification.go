// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesinstancesworkspaceinstance


type WorkspacesinstancesWorkspaceInstanceManagedInstanceCapacityReservationSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/workspacesinstances_workspace_instance#capacity_reservation_preference WorkspacesinstancesWorkspaceInstance#capacity_reservation_preference}.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/workspacesinstances_workspace_instance#capacity_reservation_target WorkspacesinstancesWorkspaceInstance#capacity_reservation_target}.
	CapacityReservationTarget *WorkspacesinstancesWorkspaceInstanceManagedInstanceCapacityReservationSpecificationCapacityReservationTarget `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

