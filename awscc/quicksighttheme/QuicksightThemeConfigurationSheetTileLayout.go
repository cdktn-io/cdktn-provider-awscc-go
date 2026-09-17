// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationSheetTileLayout struct {
	// <p>The display options for gutter spacing between tiles on a sheet.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_theme#gutter QuicksightTheme#gutter}
	Gutter *QuicksightThemeConfigurationSheetTileLayoutGutter `field:"optional" json:"gutter" yaml:"gutter"`
	// <p>The display options for margins around the outside edge of sheets.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_theme#margin QuicksightTheme#margin}
	Margin *QuicksightThemeConfigurationSheetTileLayoutMargin `field:"optional" json:"margin" yaml:"margin"`
}

