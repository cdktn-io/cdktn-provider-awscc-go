// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionsContainerDistributionConfiguration struct {
	// Tags that are attached to the container distribution configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_distribution_configuration#container_tags ImagebuilderDistributionConfiguration#container_tags}
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// The description of the container distribution configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_distribution_configuration#description ImagebuilderDistributionConfiguration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The destination repository for the container distribution configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_distribution_configuration#target_repository ImagebuilderDistributionConfiguration#target_repository}
	TargetRepository *ImagebuilderDistributionConfigurationDistributionsContainerDistributionConfigurationTargetRepository `field:"optional" json:"targetRepository" yaml:"targetRepository"`
}

