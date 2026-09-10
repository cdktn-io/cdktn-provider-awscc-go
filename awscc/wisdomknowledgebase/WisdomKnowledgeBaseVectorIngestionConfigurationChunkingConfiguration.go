// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomknowledgebase


type WisdomKnowledgeBaseVectorIngestionConfigurationChunkingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_knowledge_base#chunking_strategy WisdomKnowledgeBase#chunking_strategy}.
	ChunkingStrategy *string `field:"optional" json:"chunkingStrategy" yaml:"chunkingStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_knowledge_base#fixed_size_chunking_configuration WisdomKnowledgeBase#fixed_size_chunking_configuration}.
	FixedSizeChunkingConfiguration *WisdomKnowledgeBaseVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfiguration `field:"optional" json:"fixedSizeChunkingConfiguration" yaml:"fixedSizeChunkingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_knowledge_base#hierarchical_chunking_configuration WisdomKnowledgeBase#hierarchical_chunking_configuration}.
	HierarchicalChunkingConfiguration *WisdomKnowledgeBaseVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfiguration `field:"optional" json:"hierarchicalChunkingConfiguration" yaml:"hierarchicalChunkingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_knowledge_base#semantic_chunking_configuration WisdomKnowledgeBase#semantic_chunking_configuration}.
	SemanticChunkingConfiguration *WisdomKnowledgeBaseVectorIngestionConfigurationChunkingConfigurationSemanticChunkingConfiguration `field:"optional" json:"semanticChunkingConfiguration" yaml:"semanticChunkingConfiguration"`
}

