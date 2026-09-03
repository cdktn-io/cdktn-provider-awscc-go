// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecapacityprovider


type BedrockagentcoreCapacityProviderComputeConfiguration struct {
	// Configuration for EC2-based capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_capacity_provider#ec_2_configuration BedrockagentcoreCapacityProvider#ec_2_configuration}
	Ec2Configuration *BedrockagentcoreCapacityProviderComputeConfigurationEc2Configuration `field:"required" json:"ec2Configuration" yaml:"ec2Configuration"`
}

