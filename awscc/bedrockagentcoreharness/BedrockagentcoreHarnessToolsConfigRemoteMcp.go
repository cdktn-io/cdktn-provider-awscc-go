// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessToolsConfigRemoteMcp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_harness#headers BedrockagentcoreHarness#headers}.
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_harness#url BedrockagentcoreHarness#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

