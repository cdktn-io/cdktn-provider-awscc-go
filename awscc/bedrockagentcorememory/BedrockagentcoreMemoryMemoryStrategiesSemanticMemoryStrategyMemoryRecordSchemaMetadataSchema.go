// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryMemoryStrategiesSemanticMemoryStrategyMemoryRecordSchemaMetadataSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_memory#extraction_config BedrockagentcoreMemory#extraction_config}.
	ExtractionConfig *BedrockagentcoreMemoryMemoryStrategiesSemanticMemoryStrategyMemoryRecordSchemaMetadataSchemaExtractionConfig `field:"optional" json:"extractionConfig" yaml:"extractionConfig"`
	// Specifies whether the metadata value is extracted by the LLM or passed through deterministically from the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_memory#extraction_type BedrockagentcoreMemory#extraction_type}
	ExtractionType *string `field:"optional" json:"extractionType" yaml:"extractionType"`
	// Key name for metadata fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_memory#key BedrockagentcoreMemory#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Supported data types for metadata values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_memory#type BedrockagentcoreMemory#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

