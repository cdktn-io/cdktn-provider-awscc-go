// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#relevance_score BedrockagentcoreHarness#relevance_score}.
	RelevanceScore *float64 `field:"optional" json:"relevanceScore" yaml:"relevanceScore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#strategy_id BedrockagentcoreHarness#strategy_id}.
	StrategyId *string `field:"optional" json:"strategyId" yaml:"strategyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_harness#top_k BedrockagentcoreHarness#top_k}.
	TopK *float64 `field:"optional" json:"topK" yaml:"topK"`
}

