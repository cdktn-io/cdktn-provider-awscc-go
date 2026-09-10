// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectworkspace


type ConnectWorkspaceThemeDarkPaletteNavigation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_workspace#background ConnectWorkspace#background}.
	Background *string `field:"optional" json:"background" yaml:"background"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_workspace#invert_actions_colors ConnectWorkspace#invert_actions_colors}.
	InvertActionsColors interface{} `field:"optional" json:"invertActionsColors" yaml:"invertActionsColors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_workspace#text ConnectWorkspace#text}.
	Text *string `field:"optional" json:"text" yaml:"text"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_workspace#text_active ConnectWorkspace#text_active}.
	TextActive *string `field:"optional" json:"textActive" yaml:"textActive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_workspace#text_background_active ConnectWorkspace#text_background_active}.
	TextBackgroundActive *string `field:"optional" json:"textBackgroundActive" yaml:"textBackgroundActive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_workspace#text_background_hover ConnectWorkspace#text_background_hover}.
	TextBackgroundHover *string `field:"optional" json:"textBackgroundHover" yaml:"textBackgroundHover"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_workspace#text_hover ConnectWorkspace#text_hover}.
	TextHover *string `field:"optional" json:"textHover" yaml:"textHover"`
}

