// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfig struct {
	// Definition for the metadata schema entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_memory#definition BedrockagentcoreMemory#definition}
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// LLM extraction instruction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_memory#llm_extraction_instruction BedrockagentcoreMemory#llm_extraction_instruction}
	LlmExtractionInstruction *string `field:"optional" json:"llmExtractionInstruction" yaml:"llmExtractionInstruction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_memory#validation BedrockagentcoreMemory#validation}.
	Validation *BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidation `field:"optional" json:"validation" yaml:"validation"`
}

