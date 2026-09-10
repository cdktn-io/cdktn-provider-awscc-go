// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeNetworkConfiguration struct {
	// Network mode configuration type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_runtime#network_mode BedrockagentcoreRuntime#network_mode}
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Network mode configuration for VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_runtime#network_mode_config BedrockagentcoreRuntime#network_mode_config}
	NetworkModeConfig *BedrockagentcoreRuntimeNetworkConfigurationNetworkModeConfig `field:"optional" json:"networkModeConfig" yaml:"networkModeConfig"`
}

