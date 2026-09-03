// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecapacityprovider


type BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationVolumes struct {
	// Configuration for an EBS-backed persistent volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_capacity_provider#ebs_configuration BedrockagentcoreCapacityProvider#ebs_configuration}
	EbsConfiguration *BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationVolumesEbsConfiguration `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
}

