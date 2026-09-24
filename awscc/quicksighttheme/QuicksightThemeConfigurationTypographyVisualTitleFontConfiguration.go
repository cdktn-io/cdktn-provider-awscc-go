// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationTypographyVisualTitleFontConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#font_configuration QuicksightTheme#font_configuration}.
	FontConfiguration *QuicksightThemeConfigurationTypographyVisualTitleFontConfigurationFontConfiguration `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#text_alignment QuicksightTheme#text_alignment}.
	TextAlignment *string `field:"optional" json:"textAlignment" yaml:"textAlignment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#text_transform QuicksightTheme#text_transform}.
	TextTransform *string `field:"optional" json:"textTransform" yaml:"textTransform"`
}

