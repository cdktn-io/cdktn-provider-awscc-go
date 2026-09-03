// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecapacityprovider


type BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationVpcConfiguration struct {
	// The IDs of the security groups to associate with the instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_capacity_provider#security_groups BedrockagentcoreCapacityProvider#security_groups}
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The IDs of the subnets in which to launch instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_capacity_provider#subnets BedrockagentcoreCapacityProvider#subnets}
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

