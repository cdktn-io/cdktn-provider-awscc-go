// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagerecipe

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ImagebuilderImageRecipeConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the image recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#name ImagebuilderImageRecipe#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The parent image of the image recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#parent_image ImagebuilderImageRecipe#parent_image}
	ParentImage *string `field:"required" json:"parentImage" yaml:"parentImage"`
	// The version of the image recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#version ImagebuilderImageRecipe#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Specify additional settings and launch scripts for your build instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#additional_instance_configuration ImagebuilderImageRecipe#additional_instance_configuration}
	AdditionalInstanceConfiguration *ImagebuilderImageRecipeAdditionalInstanceConfiguration `field:"optional" json:"additionalInstanceConfiguration" yaml:"additionalInstanceConfiguration"`
	// The tags to apply to the AMI created by this image recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#ami_tags ImagebuilderImageRecipe#ami_tags}
	AmiTags *map[string]*string `field:"optional" json:"amiTags" yaml:"amiTags"`
	// The AMI watermark names to attach to the output AMI from this recipe.
	//
	// AMI watermarks are lineage markers that automatically propagate to derivative AMIs when the source AMI is copied or distributed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#ami_watermarks ImagebuilderImageRecipe#ami_watermarks}
	AmiWatermarks *[]*string `field:"optional" json:"amiWatermarks" yaml:"amiWatermarks"`
	// The block device mappings to apply when creating images from this recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#block_device_mappings ImagebuilderImageRecipe#block_device_mappings}
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// The components of the image recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#components ImagebuilderImageRecipe#components}
	Components interface{} `field:"optional" json:"components" yaml:"components"`
	// The description of the image recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#description ImagebuilderImageRecipe#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags of the image recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#tags ImagebuilderImageRecipe#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The working directory to be used during build and test workflows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#working_directory ImagebuilderImageRecipe#working_directory}
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

