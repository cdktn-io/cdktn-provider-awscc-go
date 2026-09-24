// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfiguration struct {
	// Details about how to chunk the documents in the data source.
	//
	// A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#chunking_configuration BedrockDataSource#chunking_configuration}
	ChunkingConfiguration *BedrockDataSourceVectorIngestionConfigurationChunkingConfiguration `field:"optional" json:"chunkingConfiguration" yaml:"chunkingConfiguration"`
	// Additional Enrichment Configuration for example when using GraphRag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#context_enrichment_configuration BedrockDataSource#context_enrichment_configuration}
	ContextEnrichmentConfiguration *BedrockDataSourceVectorIngestionConfigurationContextEnrichmentConfiguration `field:"optional" json:"contextEnrichmentConfiguration" yaml:"contextEnrichmentConfiguration"`
	// Settings for customizing steps in the data source content ingestion pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#custom_transformation_configuration BedrockDataSource#custom_transformation_configuration}
	CustomTransformationConfiguration *BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfiguration `field:"optional" json:"customTransformationConfiguration" yaml:"customTransformationConfiguration"`
	// Settings for parsing document contents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#parsing_configuration BedrockDataSource#parsing_configuration}
	ParsingConfiguration *BedrockDataSourceVectorIngestionConfigurationParsingConfiguration `field:"optional" json:"parsingConfiguration" yaml:"parsingConfiguration"`
}

