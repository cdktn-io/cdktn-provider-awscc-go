// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagerecipe


type ImagebuilderImageRecipeComponentsParameters struct {
	// The name of the component parameter to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/imagebuilder_image_recipe#name ImagebuilderImageRecipe#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Sets the value for the named component parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/imagebuilder_image_recipe#value ImagebuilderImageRecipe#value}
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

