// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationSheetBackground struct {
	// String to encapsulate the most generic way Color can be formatted (words, hexStrings etc).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#color QuicksightTheme#color}
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#gradient QuicksightTheme#gradient}.
	Gradient *string `field:"optional" json:"gradient" yaml:"gradient"`
}

