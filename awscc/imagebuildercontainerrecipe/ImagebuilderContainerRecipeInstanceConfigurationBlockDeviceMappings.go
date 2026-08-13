// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuildercontainerrecipe


type ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappings struct {
	// The device to which these mappings apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/imagebuilder_container_recipe#device_name ImagebuilderContainerRecipe#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Use to manage Amazon EBS-specific configuration for this mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/imagebuilder_container_recipe#ebs ImagebuilderContainerRecipe#ebs}
	Ebs *ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingsEbs `field:"optional" json:"ebs" yaml:"ebs"`
	// Use to remove a mapping from the parent image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/imagebuilder_container_recipe#no_device ImagebuilderContainerRecipe#no_device}
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Use to manage instance ephemeral devices.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/imagebuilder_container_recipe#virtual_name ImagebuilderContainerRecipe#virtual_name}
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

