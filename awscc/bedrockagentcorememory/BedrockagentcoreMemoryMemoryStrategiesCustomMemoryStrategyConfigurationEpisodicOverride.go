// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationEpisodicOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_memory#consolidation BedrockagentcoreMemory#consolidation}.
	Consolidation *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationEpisodicOverrideConsolidation `field:"optional" json:"consolidation" yaml:"consolidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_memory#extraction BedrockagentcoreMemory#extraction}.
	Extraction *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationEpisodicOverrideExtraction `field:"optional" json:"extraction" yaml:"extraction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_memory#reflection BedrockagentcoreMemory#reflection}.
	Reflection *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationEpisodicOverrideReflection `field:"optional" json:"reflection" yaml:"reflection"`
}

