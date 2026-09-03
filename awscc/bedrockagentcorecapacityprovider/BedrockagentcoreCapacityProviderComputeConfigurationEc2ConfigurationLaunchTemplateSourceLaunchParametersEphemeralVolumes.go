// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecapacityprovider


type BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationLaunchTemplateSourceLaunchParametersEphemeralVolumes struct {
	// The device name (for example, /dev/sdh or xvdh).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_capacity_provider#device_name BedrockagentcoreCapacityProvider#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Parameters used to automatically set up EBS volumes when the instance is launched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_capacity_provider#ebs BedrockagentcoreCapacityProvider#ebs}
	Ebs *BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationLaunchTemplateSourceLaunchParametersEphemeralVolumesEbs `field:"optional" json:"ebs" yaml:"ebs"`
	// The virtual device name (ephemeralN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_capacity_provider#virtual_name BedrockagentcoreCapacityProvider#virtual_name}
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

