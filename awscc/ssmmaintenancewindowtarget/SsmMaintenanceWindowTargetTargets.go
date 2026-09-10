// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtarget


type SsmMaintenanceWindowTargetTargets struct {
	// User-defined criteria for sending commands that target managed nodes that meet the criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_maintenance_window_target#key SsmMaintenanceWindowTarget#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// User-defined criteria that maps to Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_maintenance_window_target#values SsmMaintenanceWindowTarget#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

