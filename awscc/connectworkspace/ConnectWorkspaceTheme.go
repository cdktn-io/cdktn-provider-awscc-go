// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectworkspace


type ConnectWorkspaceTheme struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_workspace#dark ConnectWorkspace#dark}.
	Dark *ConnectWorkspaceThemeDark `field:"optional" json:"dark" yaml:"dark"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_workspace#light ConnectWorkspace#light}.
	Light *ConnectWorkspaceThemeLight `field:"optional" json:"light" yaml:"light"`
}

