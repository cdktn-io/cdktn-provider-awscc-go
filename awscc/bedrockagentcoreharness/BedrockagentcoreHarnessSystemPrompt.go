// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessSystemPrompt struct {
	// The text content of the system prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_harness#text BedrockagentcoreHarness#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}

