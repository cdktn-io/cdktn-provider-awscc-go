// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessTruncationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_harness#sliding_window BedrockagentcoreHarness#sliding_window}.
	SlidingWindow *BedrockagentcoreHarnessTruncationConfigSlidingWindow `field:"optional" json:"slidingWindow" yaml:"slidingWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_harness#summarization BedrockagentcoreHarness#summarization}.
	Summarization *BedrockagentcoreHarnessTruncationConfigSummarization `field:"optional" json:"summarization" yaml:"summarization"`
}

