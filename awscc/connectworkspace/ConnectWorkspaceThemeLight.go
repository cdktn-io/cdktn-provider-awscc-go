// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectworkspace


type ConnectWorkspaceThemeLight struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#palette ConnectWorkspace#palette}.
	Palette *ConnectWorkspaceThemeLightPalette `field:"optional" json:"palette" yaml:"palette"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#typography ConnectWorkspace#typography}.
	Typography *ConnectWorkspaceThemeLightTypography `field:"optional" json:"typography" yaml:"typography"`
}

