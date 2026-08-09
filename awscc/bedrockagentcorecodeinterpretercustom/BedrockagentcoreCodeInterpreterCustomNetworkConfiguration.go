// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecodeinterpretercustom


type BedrockagentcoreCodeInterpreterCustomNetworkConfiguration struct {
	// Network modes supported by code interpreter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_code_interpreter_custom#network_mode BedrockagentcoreCodeInterpreterCustom#network_mode}
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Network mode configuration for VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_code_interpreter_custom#vpc_config BedrockagentcoreCodeInterpreterCustom#vpc_config}
	VpcConfig *BedrockagentcoreCodeInterpreterCustomNetworkConfigurationVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

