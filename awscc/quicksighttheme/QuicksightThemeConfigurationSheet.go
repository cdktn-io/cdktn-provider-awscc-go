// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationSheet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#background QuicksightTheme#background}.
	Background *QuicksightThemeConfigurationSheetBackground `field:"optional" json:"background" yaml:"background"`
	// <p>Display options related to tiles on a sheet.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#tile QuicksightTheme#tile}
	Tile *QuicksightThemeConfigurationSheetTile `field:"optional" json:"tile" yaml:"tile"`
	// <p>The display options for the layout of tiles on a sheet.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_theme#tile_layout QuicksightTheme#tile_layout}
	TileLayout *QuicksightThemeConfigurationSheetTileLayout `field:"optional" json:"tileLayout" yaml:"tileLayout"`
}

