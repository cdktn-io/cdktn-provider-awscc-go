// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationUiColorPalette struct {
	// <p>This color is that applies to selected states and buttons.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#accent QuicksightTheme#accent}
	Accent *string `field:"optional" json:"accent" yaml:"accent"`
	// <p>The foreground color that applies to any text or other elements that appear over the             accent color.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#accent_foreground QuicksightTheme#accent_foreground}
	AccentForeground *string `field:"optional" json:"accentForeground" yaml:"accentForeground"`
	// <p>The color that applies to error messages.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#danger QuicksightTheme#danger}
	Danger *string `field:"optional" json:"danger" yaml:"danger"`
	// <p>The foreground color that applies to any text or other elements that appear over the             error color.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#danger_foreground QuicksightTheme#danger_foreground}
	DangerForeground *string `field:"optional" json:"dangerForeground" yaml:"dangerForeground"`
	// <p>The color that applies to the names of fields that are identified as             dimensions.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#dimension QuicksightTheme#dimension}
	Dimension *string `field:"optional" json:"dimension" yaml:"dimension"`
	// <p>The foreground color that applies to any text or other elements that appear over the             dimension color.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#dimension_foreground QuicksightTheme#dimension_foreground}
	DimensionForeground *string `field:"optional" json:"dimensionForeground" yaml:"dimensionForeground"`
	// <p>The color that applies to the names of fields that are identified as measures.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#measure QuicksightTheme#measure}
	Measure *string `field:"optional" json:"measure" yaml:"measure"`
	// <p>The foreground color that applies to any text or other elements that appear over the             measure color.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#measure_foreground QuicksightTheme#measure_foreground}
	MeasureForeground *string `field:"optional" json:"measureForeground" yaml:"measureForeground"`
	// <p>The background color that applies to visuals and other high emphasis UI.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#primary_background QuicksightTheme#primary_background}
	PrimaryBackground *string `field:"optional" json:"primaryBackground" yaml:"primaryBackground"`
	// <p>The color of text and other foreground elements that appear over the primary             background regions, such as grid lines, borders, table banding, icons, and so on.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#primary_foreground QuicksightTheme#primary_foreground}
	PrimaryForeground *string `field:"optional" json:"primaryForeground" yaml:"primaryForeground"`
	// <p>The background color that applies to the sheet background and sheet controls.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#secondary_background QuicksightTheme#secondary_background}
	SecondaryBackground *string `field:"optional" json:"secondaryBackground" yaml:"secondaryBackground"`
	// <p>The foreground color that applies to any sheet title, sheet control text, or UI that             appears over the secondary background.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#secondary_foreground QuicksightTheme#secondary_foreground}
	SecondaryForeground *string `field:"optional" json:"secondaryForeground" yaml:"secondaryForeground"`
	// <p>The color that applies to success messages, for example the check mark for a             successful download.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#success QuicksightTheme#success}
	Success *string `field:"optional" json:"success" yaml:"success"`
	// <p>The foreground color that applies to any text or other elements that appear over the             success color.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#success_foreground QuicksightTheme#success_foreground}
	SuccessForeground *string `field:"optional" json:"successForeground" yaml:"successForeground"`
	// <p>This color that applies to warning and informational messages.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#warning QuicksightTheme#warning}
	Warning *string `field:"optional" json:"warning" yaml:"warning"`
	// <p>The foreground color that applies to any text or other elements that appear over the             warning color.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_theme#warning_foreground QuicksightTheme#warning_foreground}
	WarningForeground *string `field:"optional" json:"warningForeground" yaml:"warningForeground"`
}

