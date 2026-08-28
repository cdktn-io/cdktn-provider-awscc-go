// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagerecipe


type ImagebuilderImageRecipeBlockDeviceMappings struct {
	// The device to which these mappings apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#device_name ImagebuilderImageRecipe#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Use to manage Amazon EBS-specific configuration for this mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#ebs ImagebuilderImageRecipe#ebs}
	Ebs *ImagebuilderImageRecipeBlockDeviceMappingsEbs `field:"optional" json:"ebs" yaml:"ebs"`
	// Use to remove a mapping from the parent image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#no_device ImagebuilderImageRecipe#no_device}
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Use to manage instance ephemeral devices.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_image_recipe#virtual_name ImagebuilderImageRecipe#virtual_name}
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

