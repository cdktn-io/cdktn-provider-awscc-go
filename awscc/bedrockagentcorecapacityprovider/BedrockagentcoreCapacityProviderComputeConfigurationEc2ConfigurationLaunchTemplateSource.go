// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecapacityprovider


type BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationLaunchTemplateSource struct {
	// Parameters for launching EC2 instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_capacity_provider#launch_parameters BedrockagentcoreCapacityProvider#launch_parameters}
	LaunchParameters *BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationLaunchTemplateSourceLaunchParameters `field:"required" json:"launchParameters" yaml:"launchParameters"`
}

