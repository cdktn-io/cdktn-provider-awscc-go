// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectoryStreamingPropertiesUserSettings struct {
	// Indicates the type of action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#action WorkspacesDirectory#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Indicates the maximum character length for the specified user setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#maximum_length WorkspacesDirectory#maximum_length}
	MaximumLength *float64 `field:"optional" json:"maximumLength" yaml:"maximumLength"`
	// Indicates if the setting is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#permission WorkspacesDirectory#permission}
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

