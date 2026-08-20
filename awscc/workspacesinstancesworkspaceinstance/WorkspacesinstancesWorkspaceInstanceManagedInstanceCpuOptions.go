// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesinstancesworkspaceinstance


type WorkspacesinstancesWorkspaceInstanceManagedInstanceCpuOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/workspacesinstances_workspace_instance#core_count WorkspacesinstancesWorkspaceInstance#core_count}.
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/workspacesinstances_workspace_instance#threads_per_core WorkspacesinstancesWorkspaceInstance#threads_per_core}.
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

