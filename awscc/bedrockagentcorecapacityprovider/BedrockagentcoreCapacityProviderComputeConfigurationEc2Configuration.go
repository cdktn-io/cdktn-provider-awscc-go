// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecapacityprovider


type BedrockagentcoreCapacityProviderComputeConfigurationEc2Configuration struct {
	// How the launch template is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_capacity_provider#launch_template_source BedrockagentcoreCapacityProvider#launch_template_source}
	LaunchTemplateSource *BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationLaunchTemplateSource `field:"required" json:"launchTemplateSource" yaml:"launchTemplateSource"`
	// VPC configuration for launching EC2 instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_capacity_provider#vpc_configuration BedrockagentcoreCapacityProvider#vpc_configuration}
	VpcConfiguration *BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationVpcConfiguration `field:"required" json:"vpcConfiguration" yaml:"vpcConfiguration"`
	// Configuration for managing the lifecycle of instances in a capacity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_capacity_provider#lifecycle_configuration BedrockagentcoreCapacityProvider#lifecycle_configuration}
	LifecycleConfiguration *BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationLifecycleConfiguration `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// Customer-facing configuration for the (service-managed) root volume.
	//
	// The service provisions the root volume at its own AMI size estimate plus FreeSpaceGiB, and pins the visible free space to FreeSpaceGiB with a filler file, so the space you are guaranteed does not change as the underlying AMI grows. The device name and the delete-on-termination behavior are service-owned and are not configurable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_capacity_provider#root_volume BedrockagentcoreCapacityProvider#root_volume}
	RootVolume *BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationRootVolume `field:"optional" json:"rootVolume" yaml:"rootVolume"`
	// Named persistent EBS volumes for this capacity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_capacity_provider#volumes BedrockagentcoreCapacityProvider#volumes}
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

