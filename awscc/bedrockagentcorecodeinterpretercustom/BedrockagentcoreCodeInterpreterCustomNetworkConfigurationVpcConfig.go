// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecodeinterpretercustom


type BedrockagentcoreCodeInterpreterCustomNetworkConfigurationVpcConfig struct {
	// Security groups for VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_code_interpreter_custom#security_groups BedrockagentcoreCodeInterpreterCustom#security_groups}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Subnets for VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_code_interpreter_custom#subnets BedrockagentcoreCodeInterpreterCustom#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

