// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightbrand


type QuicksightBrandBrandDefinitionApplicationTheme struct {
	// The color palette.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/quicksight_brand#brand_color_palette QuicksightBrand#brand_color_palette}
	BrandColorPalette *QuicksightBrandBrandDefinitionApplicationThemeBrandColorPalette `field:"optional" json:"brandColorPalette" yaml:"brandColorPalette"`
	// The element style.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/quicksight_brand#brand_element_style QuicksightBrand#brand_element_style}
	BrandElementStyle *QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyle `field:"optional" json:"brandElementStyle" yaml:"brandElementStyle"`
}

