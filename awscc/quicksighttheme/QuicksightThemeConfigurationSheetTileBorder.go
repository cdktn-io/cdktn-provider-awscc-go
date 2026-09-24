// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationSheetTileBorder struct {
	// String to encapsulate the most generic way Color can be formatted (words, hexStrings etc).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#color QuicksightTheme#color}
	Color *string `field:"optional" json:"color" yaml:"color"`
	// <p>The option to enable display of borders for visuals.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#show QuicksightTheme#show}
	Show interface{} `field:"optional" json:"show" yaml:"show"`
	// String to encapsulate the most generic way Width can be formatted with whatever units (px, em etc).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#width QuicksightTheme#width}
	Width *string `field:"optional" json:"width" yaml:"width"`
}

