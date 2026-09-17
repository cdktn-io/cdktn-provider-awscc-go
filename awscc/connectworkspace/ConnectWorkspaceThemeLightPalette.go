// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectworkspace


type ConnectWorkspaceThemeLightPalette struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_workspace#canvas ConnectWorkspace#canvas}.
	Canvas *ConnectWorkspaceThemeLightPaletteCanvas `field:"optional" json:"canvas" yaml:"canvas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_workspace#header ConnectWorkspace#header}.
	Header *ConnectWorkspaceThemeLightPaletteHeader `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_workspace#navigation ConnectWorkspace#navigation}.
	Navigation *ConnectWorkspaceThemeLightPaletteNavigation `field:"optional" json:"navigation" yaml:"navigation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_workspace#primary ConnectWorkspace#primary}.
	Primary *ConnectWorkspaceThemeLightPalettePrimary `field:"optional" json:"primary" yaml:"primary"`
}

