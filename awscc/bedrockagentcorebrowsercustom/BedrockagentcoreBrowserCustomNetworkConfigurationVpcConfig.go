// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorebrowsercustom


type BedrockagentcoreBrowserCustomNetworkConfigurationVpcConfig struct {
	// Security groups for VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_browser_custom#security_groups BedrockagentcoreBrowserCustom#security_groups}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Subnets for VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_browser_custom#subnets BedrockagentcoreBrowserCustom#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

