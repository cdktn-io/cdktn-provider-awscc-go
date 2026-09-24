// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationSheetTile struct {
	// String to encapsulate the most generic way Color can be formatted (words, hexStrings etc).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#background_color QuicksightTheme#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// <p>The display options for tile borders for visuals.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#border QuicksightTheme#border}
	Border *QuicksightThemeConfigurationSheetTileBorder `field:"optional" json:"border" yaml:"border"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#border_radius QuicksightTheme#border_radius}.
	BorderRadius *string `field:"optional" json:"borderRadius" yaml:"borderRadius"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#padding QuicksightTheme#padding}.
	Padding *string `field:"optional" json:"padding" yaml:"padding"`
}

