// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightbrand


type QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyle struct {
	// The palette colors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#contextual_navbar QuicksightBrand#contextual_navbar}
	ContextualNavbar *QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleContextualNavbar `field:"optional" json:"contextualNavbar" yaml:"contextualNavbar"`
	// The palette colors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_brand#global_navbar QuicksightBrand#global_navbar}
	GlobalNavbar *QuicksightBrandBrandDefinitionApplicationThemeBrandElementStyleNavbarStyleGlobalNavbar `field:"optional" json:"globalNavbar" yaml:"globalNavbar"`
}

