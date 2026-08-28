// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfiguration struct {
	// <p>The theme colors that are used for data colors in charts.
	//
	// The colors description is a
	//             hexadecimal color code that consists of six alphanumerical characters, prefixed with
	//                 <code>#</code>, for example #37BFF5. </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_theme#data_color_palette QuicksightTheme#data_color_palette}
	DataColorPalette *QuicksightThemeConfigurationDataColorPalette `field:"optional" json:"dataColorPalette" yaml:"dataColorPalette"`
	// <p>The theme display options for sheets. </p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_theme#sheet QuicksightTheme#sheet}
	Sheet *QuicksightThemeConfigurationSheet `field:"optional" json:"sheet" yaml:"sheet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_theme#typography QuicksightTheme#typography}.
	Typography *QuicksightThemeConfigurationTypography `field:"optional" json:"typography" yaml:"typography"`
	// <p>The theme colors that apply to UI and to charts, excluding data colors.
	//
	// The colors
	//             description is a hexadecimal color code that consists of six alphanumerical characters,
	//             prefixed with <code>#</code>, for example #37BFF5. For more information, see <a href="https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html">Using Themes in Amazon QuickSight</a> in the <i>Amazon QuickSight User
	//                 Guide.</i>
	//          </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_theme#ui_color_palette QuicksightTheme#ui_color_palette}
	UiColorPalette *QuicksightThemeConfigurationUiColorPalette `field:"optional" json:"uiColorPalette" yaml:"uiColorPalette"`
}

