// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindow


type SsmMaintenanceWindowTags struct {
	// The name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window#key SsmMaintenanceWindow#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_maintenance_window#value SsmMaintenanceWindow#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

