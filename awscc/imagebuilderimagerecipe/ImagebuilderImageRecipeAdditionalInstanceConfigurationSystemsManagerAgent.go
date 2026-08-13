// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagerecipe


type ImagebuilderImageRecipeAdditionalInstanceConfigurationSystemsManagerAgent struct {
	// Controls whether the SSM agent is removed from your final build image, prior to creating the new AMI.
	//
	// If this is set to true, then the agent is removed from the final image. If it's set to false, then the agent is left in, so that it is included in the new AMI. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/imagebuilder_image_recipe#uninstall_after_build ImagebuilderImageRecipe#uninstall_after_build}
	UninstallAfterBuild interface{} `field:"optional" json:"uninstallAfterBuild" yaml:"uninstallAfterBuild"`
}

