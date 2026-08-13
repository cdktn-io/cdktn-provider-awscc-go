// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceswebusersettings


type WorkspaceswebUserSettingsBrandingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/workspacesweb_user_settings#color_theme WorkspaceswebUserSettings#color_theme}.
	ColorTheme *string `field:"optional" json:"colorTheme" yaml:"colorTheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/workspacesweb_user_settings#favicon WorkspaceswebUserSettings#favicon}.
	Favicon *string `field:"optional" json:"favicon" yaml:"favicon"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/workspacesweb_user_settings#localized_strings WorkspaceswebUserSettings#localized_strings}.
	LocalizedStrings interface{} `field:"optional" json:"localizedStrings" yaml:"localizedStrings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/workspacesweb_user_settings#logo WorkspaceswebUserSettings#logo}.
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/workspacesweb_user_settings#terms_of_service WorkspaceswebUserSettings#terms_of_service}.
	TermsOfService *string `field:"optional" json:"termsOfService" yaml:"termsOfService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/workspacesweb_user_settings#wallpaper WorkspaceswebUserSettings#wallpaper}.
	Wallpaper *string `field:"optional" json:"wallpaper" yaml:"wallpaper"`
}

