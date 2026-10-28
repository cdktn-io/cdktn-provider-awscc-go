// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightbrand


type QuicksightBrandBrandDefinition struct {
	// The application theme.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#application_theme QuicksightBrand#application_theme}
	ApplicationTheme *QuicksightBrandBrandDefinitionApplicationTheme `field:"optional" json:"applicationTheme" yaml:"applicationTheme"`
	// The name of the brand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#brand_name QuicksightBrand#brand_name}
	BrandName *string `field:"optional" json:"brandName" yaml:"brandName"`
	// The description of the brand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#description QuicksightBrand#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

