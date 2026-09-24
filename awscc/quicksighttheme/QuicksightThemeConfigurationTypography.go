// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationTypography struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#axis_label_font_configuration QuicksightTheme#axis_label_font_configuration}.
	AxisLabelFontConfiguration *QuicksightThemeConfigurationTypographyAxisLabelFontConfiguration `field:"optional" json:"axisLabelFontConfiguration" yaml:"axisLabelFontConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#axis_title_font_configuration QuicksightTheme#axis_title_font_configuration}.
	AxisTitleFontConfiguration *QuicksightThemeConfigurationTypographyAxisTitleFontConfiguration `field:"optional" json:"axisTitleFontConfiguration" yaml:"axisTitleFontConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#data_label_font_configuration QuicksightTheme#data_label_font_configuration}.
	DataLabelFontConfiguration *QuicksightThemeConfigurationTypographyDataLabelFontConfiguration `field:"optional" json:"dataLabelFontConfiguration" yaml:"dataLabelFontConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#font_families QuicksightTheme#font_families}.
	FontFamilies interface{} `field:"optional" json:"fontFamilies" yaml:"fontFamilies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#legend_title_font_configuration QuicksightTheme#legend_title_font_configuration}.
	LegendTitleFontConfiguration *QuicksightThemeConfigurationTypographyLegendTitleFontConfiguration `field:"optional" json:"legendTitleFontConfiguration" yaml:"legendTitleFontConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#legend_value_font_configuration QuicksightTheme#legend_value_font_configuration}.
	LegendValueFontConfiguration *QuicksightThemeConfigurationTypographyLegendValueFontConfiguration `field:"optional" json:"legendValueFontConfiguration" yaml:"legendValueFontConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#visual_subtitle_font_configuration QuicksightTheme#visual_subtitle_font_configuration}.
	VisualSubtitleFontConfiguration *QuicksightThemeConfigurationTypographyVisualSubtitleFontConfiguration `field:"optional" json:"visualSubtitleFontConfiguration" yaml:"visualSubtitleFontConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#visual_title_font_configuration QuicksightTheme#visual_title_font_configuration}.
	VisualTitleFontConfiguration *QuicksightThemeConfigurationTypographyVisualTitleFontConfiguration `field:"optional" json:"visualTitleFontConfiguration" yaml:"visualTitleFontConfiguration"`
}

