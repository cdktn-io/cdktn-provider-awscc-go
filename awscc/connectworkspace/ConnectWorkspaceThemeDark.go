// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectworkspace


type ConnectWorkspaceThemeDark struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_workspace#palette ConnectWorkspace#palette}.
	Palette *ConnectWorkspaceThemeDarkPalette `field:"optional" json:"palette" yaml:"palette"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_workspace#typography ConnectWorkspace#typography}.
	Typography *ConnectWorkspaceThemeDarkTypography `field:"optional" json:"typography" yaml:"typography"`
}

