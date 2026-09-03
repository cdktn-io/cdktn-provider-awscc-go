// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationChunkingConfiguration struct {
	// Knowledge base can split your source data into chunks.
	//
	// A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried. You have the following options for chunking your data. If you opt for NONE, then you may want to pre-process your files by splitting them up such that each file corresponds to a chunk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#chunking_strategy BedrockDataSource#chunking_strategy}
	ChunkingStrategy *string `field:"optional" json:"chunkingStrategy" yaml:"chunkingStrategy"`
	// Configurations for when you choose fixed-size chunking. If you set the chunkingStrategy as NONE, exclude this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#fixed_size_chunking_configuration BedrockDataSource#fixed_size_chunking_configuration}
	FixedSizeChunkingConfiguration *BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfiguration `field:"optional" json:"fixedSizeChunkingConfiguration" yaml:"fixedSizeChunkingConfiguration"`
	// Configurations for when you choose hierarchical chunking. If you set the chunkingStrategy as NONE, exclude this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#hierarchical_chunking_configuration BedrockDataSource#hierarchical_chunking_configuration}
	HierarchicalChunkingConfiguration *BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfiguration `field:"optional" json:"hierarchicalChunkingConfiguration" yaml:"hierarchicalChunkingConfiguration"`
	// Configurations for when you choose semantic chunking. If you set the chunkingStrategy as NONE, exclude this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#semantic_chunking_configuration BedrockDataSource#semantic_chunking_configuration}
	SemanticChunkingConfiguration *BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationSemanticChunkingConfiguration `field:"optional" json:"semanticChunkingConfiguration" yaml:"semanticChunkingConfiguration"`
}

