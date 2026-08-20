// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesinstancesworkspaceinstance


type WorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkInterfaces struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/workspacesinstances_workspace_instance#description WorkspacesinstancesWorkspaceInstance#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/workspacesinstances_workspace_instance#device_index WorkspacesinstancesWorkspaceInstance#device_index}.
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/workspacesinstances_workspace_instance#groups WorkspacesinstancesWorkspaceInstance#groups}.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/workspacesinstances_workspace_instance#subnet_id WorkspacesinstancesWorkspaceInstance#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

