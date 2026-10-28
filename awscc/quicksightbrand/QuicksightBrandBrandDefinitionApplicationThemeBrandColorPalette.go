// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightbrand


type QuicksightBrandBrandDefinitionApplicationThemeBrandColorPalette struct {
	// The palette colors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#accent QuicksightBrand#accent}
	Accent *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteAccent `field:"optional" json:"accent" yaml:"accent"`
	// The palette colors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#danger QuicksightBrand#danger}
	Danger *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDanger `field:"optional" json:"danger" yaml:"danger"`
	// The palette colors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#dimension QuicksightBrand#dimension}
	Dimension *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// The palette colors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#info QuicksightBrand#info}
	Info *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteInfo `field:"optional" json:"info" yaml:"info"`
	// The palette colors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#measure QuicksightBrand#measure}
	Measure *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteMeasure `field:"optional" json:"measure" yaml:"measure"`
	// The palette colors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#primary QuicksightBrand#primary}
	Primary *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPalettePrimary `field:"optional" json:"primary" yaml:"primary"`
	// The palette colors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#secondary QuicksightBrand#secondary}
	Secondary *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSecondary `field:"optional" json:"secondary" yaml:"secondary"`
	// The palette colors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#success QuicksightBrand#success}
	Success *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteSuccess `field:"optional" json:"success" yaml:"success"`
	// The palette colors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#warning QuicksightBrand#warning}
	Warning *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPaletteWarning `field:"optional" json:"warning" yaml:"warning"`
}

