// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryMemoryStrategies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_memory#custom_memory_strategy BedrockagentcoreMemory#custom_memory_strategy}.
	CustomMemoryStrategy *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategy `field:"optional" json:"customMemoryStrategy" yaml:"customMemoryStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_memory#episodic_memory_strategy BedrockagentcoreMemory#episodic_memory_strategy}.
	EpisodicMemoryStrategy *BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategy `field:"optional" json:"episodicMemoryStrategy" yaml:"episodicMemoryStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_memory#semantic_memory_strategy BedrockagentcoreMemory#semantic_memory_strategy}.
	SemanticMemoryStrategy *BedrockagentcoreMemoryMemoryStrategiesSemanticMemoryStrategy `field:"optional" json:"semanticMemoryStrategy" yaml:"semanticMemoryStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_memory#summary_memory_strategy BedrockagentcoreMemory#summary_memory_strategy}.
	SummaryMemoryStrategy *BedrockagentcoreMemoryMemoryStrategiesSummaryMemoryStrategy `field:"optional" json:"summaryMemoryStrategy" yaml:"summaryMemoryStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_memory#user_preference_memory_strategy BedrockagentcoreMemory#user_preference_memory_strategy}.
	UserPreferenceMemoryStrategy *BedrockagentcoreMemoryMemoryStrategiesUserPreferenceMemoryStrategy `field:"optional" json:"userPreferenceMemoryStrategy" yaml:"userPreferenceMemoryStrategy"`
}

