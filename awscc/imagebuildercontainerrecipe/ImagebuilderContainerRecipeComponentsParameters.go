// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuildercontainerrecipe


type ImagebuilderContainerRecipeComponentsParameters struct {
	// The name of the component parameter to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/imagebuilder_container_recipe#name ImagebuilderContainerRecipe#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Sets the value for the named component parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/imagebuilder_container_recipe#value ImagebuilderContainerRecipe#value}
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

