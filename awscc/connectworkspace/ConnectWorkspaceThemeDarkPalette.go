// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectworkspace


type ConnectWorkspaceThemeDarkPalette struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_workspace#canvas ConnectWorkspace#canvas}.
	Canvas *ConnectWorkspaceThemeDarkPaletteCanvas `field:"optional" json:"canvas" yaml:"canvas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_workspace#header ConnectWorkspace#header}.
	Header *ConnectWorkspaceThemeDarkPaletteHeader `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_workspace#navigation ConnectWorkspace#navigation}.
	Navigation *ConnectWorkspaceThemeDarkPaletteNavigation `field:"optional" json:"navigation" yaml:"navigation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_workspace#primary ConnectWorkspace#primary}.
	Primary *ConnectWorkspaceThemeDarkPalettePrimary `field:"optional" json:"primary" yaml:"primary"`
}

