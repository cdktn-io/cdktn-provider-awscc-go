// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessTruncation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_harness#config BedrockagentcoreHarness#config}.
	Config *BedrockagentcoreHarnessTruncationConfig `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_harness#strategy BedrockagentcoreHarness#strategy}.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

