// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectoryStreamingPropertiesGlobalAccelerator struct {
	// Indicates if Global Accelerator for directory is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#mode WorkspacesDirectory#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Indicates the preferred protocol for Global Accelerator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#preferred_protocol WorkspacesDirectory#preferred_protocol}
	PreferredProtocol *string `field:"optional" json:"preferredProtocol" yaml:"preferredProtocol"`
}

