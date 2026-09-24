// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagerecipe


type ImagebuilderImageRecipeAdditionalInstanceConfiguration struct {
	// Contains settings for the SSM agent on your build instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/imagebuilder_image_recipe#systems_manager_agent ImagebuilderImageRecipe#systems_manager_agent}
	SystemsManagerAgent *ImagebuilderImageRecipeAdditionalInstanceConfigurationSystemsManagerAgent `field:"optional" json:"systemsManagerAgent" yaml:"systemsManagerAgent"`
	// Use this property to provide commands or a command script to run when you launch your build instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/imagebuilder_image_recipe#user_data_override ImagebuilderImageRecipe#user_data_override}
	UserDataOverride *string `field:"optional" json:"userDataOverride" yaml:"userDataOverride"`
}

