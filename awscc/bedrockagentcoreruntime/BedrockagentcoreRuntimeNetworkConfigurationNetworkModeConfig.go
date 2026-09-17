// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeNetworkConfigurationNetworkModeConfig struct {
	// Security groups for VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#security_groups BedrockagentcoreRuntime#security_groups}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Subnets for VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#subnets BedrockagentcoreRuntime#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

