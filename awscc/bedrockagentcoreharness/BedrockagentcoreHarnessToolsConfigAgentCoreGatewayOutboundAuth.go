// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessToolsConfigAgentCoreGatewayOutboundAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_harness#aws_iam BedrockagentcoreHarness#aws_iam}.
	AwsIam *string `field:"optional" json:"awsIam" yaml:"awsIam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_harness#none BedrockagentcoreHarness#none}.
	None *string `field:"optional" json:"none" yaml:"none"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_harness#oauth BedrockagentcoreHarness#oauth}.
	Oauth *BedrockagentcoreHarnessToolsConfigAgentCoreGatewayOutboundAuthOauth `field:"optional" json:"oauth" yaml:"oauth"`
}

