// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationDataColorPalette struct {
	// <p>The hexadecimal codes for the colors.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_theme#colors QuicksightTheme#colors}
	Colors *[]*string `field:"optional" json:"colors" yaml:"colors"`
	// <p>The hexadecimal code of a color that applies to charts where a lack of data is             highlighted.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_theme#empty_fill_color QuicksightTheme#empty_fill_color}
	EmptyFillColor *string `field:"optional" json:"emptyFillColor" yaml:"emptyFillColor"`
	// <p>The minimum and maximum hexadecimal codes that describe a color gradient. </p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_theme#min_max_gradient QuicksightTheme#min_max_gradient}
	MinMaxGradient *[]*string `field:"optional" json:"minMaxGradient" yaml:"minMaxGradient"`
}

