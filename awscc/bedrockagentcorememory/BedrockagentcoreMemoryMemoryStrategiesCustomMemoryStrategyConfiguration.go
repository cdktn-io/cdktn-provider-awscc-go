// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#episodic_override BedrockagentcoreMemory#episodic_override}.
	EpisodicOverride *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationEpisodicOverride `field:"optional" json:"episodicOverride" yaml:"episodicOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#self_managed_configuration BedrockagentcoreMemory#self_managed_configuration}.
	SelfManagedConfiguration *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSelfManagedConfiguration `field:"optional" json:"selfManagedConfiguration" yaml:"selfManagedConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#semantic_override BedrockagentcoreMemory#semantic_override}.
	SemanticOverride *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSemanticOverride `field:"optional" json:"semanticOverride" yaml:"semanticOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#summary_override BedrockagentcoreMemory#summary_override}.
	SummaryOverride *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSummaryOverride `field:"optional" json:"summaryOverride" yaml:"summaryOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#user_preference_override BedrockagentcoreMemory#user_preference_override}.
	UserPreferenceOverride *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationUserPreferenceOverride `field:"optional" json:"userPreferenceOverride" yaml:"userPreferenceOverride"`
}

