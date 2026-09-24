// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryMemoryStrategiesSemanticMemoryStrategy struct {
	// Creation timestamp of the memory strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#created_at BedrockagentcoreMemory#created_at}
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Description of the Memory resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#description BedrockagentcoreMemory#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#memory_record_schema BedrockagentcoreMemory#memory_record_schema}.
	MemoryRecordSchema *BedrockagentcoreMemoryMemoryStrategiesSemanticMemoryStrategyMemoryRecordSchema `field:"optional" json:"memoryRecordSchema" yaml:"memoryRecordSchema"`
	// Name of the Memory resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#name BedrockagentcoreMemory#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// List of namespaces for memory strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#namespaces BedrockagentcoreMemory#namespaces}
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// List of namespaces for memory strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#namespace_templates BedrockagentcoreMemory#namespace_templates}
	NamespaceTemplates *[]*string `field:"optional" json:"namespaceTemplates" yaml:"namespaceTemplates"`
	// Status of the memory strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#status BedrockagentcoreMemory#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Unique identifier for the memory strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#strategy_id BedrockagentcoreMemory#strategy_id}
	StrategyId *string `field:"optional" json:"strategyId" yaml:"strategyId"`
	// Type of memory strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#type BedrockagentcoreMemory#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Last update timestamp of the memory strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_memory#updated_at BedrockagentcoreMemory#updated_at}
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

