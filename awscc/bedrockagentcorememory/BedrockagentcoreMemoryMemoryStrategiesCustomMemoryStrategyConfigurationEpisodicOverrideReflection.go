// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationEpisodicOverrideReflection struct {
	// Text prompt for model instructions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_memory#append_to_prompt BedrockagentcoreMemory#append_to_prompt}
	AppendToPrompt *string `field:"optional" json:"appendToPrompt" yaml:"appendToPrompt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_memory#memory_record_schema BedrockagentcoreMemory#memory_record_schema}.
	MemoryRecordSchema *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationEpisodicOverrideReflectionMemoryRecordSchema `field:"optional" json:"memoryRecordSchema" yaml:"memoryRecordSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_memory#model_id BedrockagentcoreMemory#model_id}.
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// List of namespaces for memory strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_memory#namespaces BedrockagentcoreMemory#namespaces}
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// List of namespaces for memory strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_memory#namespace_templates BedrockagentcoreMemory#namespace_templates}
	NamespaceTemplates *[]*string `field:"optional" json:"namespaceTemplates" yaml:"namespaceTemplates"`
}

