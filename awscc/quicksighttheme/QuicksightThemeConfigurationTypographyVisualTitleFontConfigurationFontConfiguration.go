// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationTypographyVisualTitleFontConfigurationFontConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#font_color QuicksightTheme#font_color}.
	FontColor *string `field:"optional" json:"fontColor" yaml:"fontColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#font_decoration QuicksightTheme#font_decoration}.
	FontDecoration *string `field:"optional" json:"fontDecoration" yaml:"fontDecoration"`
	// <p>The font family that you want to use.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#font_family QuicksightTheme#font_family}
	FontFamily *string `field:"optional" json:"fontFamily" yaml:"fontFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#font_size QuicksightTheme#font_size}.
	FontSize *QuicksightThemeConfigurationTypographyVisualTitleFontConfigurationFontConfigurationFontSize `field:"optional" json:"fontSize" yaml:"fontSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#font_style QuicksightTheme#font_style}.
	FontStyle *string `field:"optional" json:"fontStyle" yaml:"fontStyle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#font_weight QuicksightTheme#font_weight}.
	FontWeight *QuicksightThemeConfigurationTypographyVisualTitleFontConfigurationFontConfigurationFontWeight `field:"optional" json:"fontWeight" yaml:"fontWeight"`
}

