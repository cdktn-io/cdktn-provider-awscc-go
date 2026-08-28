// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributions struct {
	// region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_distribution_configuration#region ImagebuilderDistributionConfiguration#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The specific AMI settings (for example, launch permissions, AMI tags).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_distribution_configuration#ami_distribution_configuration ImagebuilderDistributionConfiguration#ami_distribution_configuration}
	AmiDistributionConfiguration *ImagebuilderDistributionConfigurationDistributionsAmiDistributionConfiguration `field:"optional" json:"amiDistributionConfiguration" yaml:"amiDistributionConfiguration"`
	// Container distribution settings for encryption, licensing, and sharing in a specific Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_distribution_configuration#container_distribution_configuration ImagebuilderDistributionConfiguration#container_distribution_configuration}
	ContainerDistributionConfiguration *ImagebuilderDistributionConfigurationDistributionsContainerDistributionConfiguration `field:"optional" json:"containerDistributionConfiguration" yaml:"containerDistributionConfiguration"`
	// The Windows faster-launching configurations to use for AMI distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_distribution_configuration#fast_launch_configurations ImagebuilderDistributionConfiguration#fast_launch_configurations}
	FastLaunchConfigurations interface{} `field:"optional" json:"fastLaunchConfigurations" yaml:"fastLaunchConfigurations"`
	// A group of launchTemplateConfiguration settings that apply to image distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_distribution_configuration#launch_template_configurations ImagebuilderDistributionConfiguration#launch_template_configurations}
	LaunchTemplateConfigurations interface{} `field:"optional" json:"launchTemplateConfigurations" yaml:"launchTemplateConfigurations"`
	// The License Manager Configuration to associate with the AMI in the specified Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_distribution_configuration#license_configuration_arns ImagebuilderDistributionConfiguration#license_configuration_arns}
	LicenseConfigurationArns *[]*string `field:"optional" json:"licenseConfigurationArns" yaml:"licenseConfigurationArns"`
	// The SSM parameter configurations to use for AMI distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_distribution_configuration#ssm_parameter_configurations ImagebuilderDistributionConfiguration#ssm_parameter_configurations}
	SsmParameterConfigurations interface{} `field:"optional" json:"ssmParameterConfigurations" yaml:"ssmParameterConfigurations"`
}

