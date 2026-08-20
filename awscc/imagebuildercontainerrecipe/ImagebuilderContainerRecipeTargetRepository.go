// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuildercontainerrecipe


type ImagebuilderContainerRecipeTargetRepository struct {
	// The name of the container repository where the output container image is stored.
	//
	// This name is prefixed by the repository location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/imagebuilder_container_recipe#repository_name ImagebuilderContainerRecipe#repository_name}
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// Specifies the service in which this image was registered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/imagebuilder_container_recipe#service ImagebuilderContainerRecipe#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

