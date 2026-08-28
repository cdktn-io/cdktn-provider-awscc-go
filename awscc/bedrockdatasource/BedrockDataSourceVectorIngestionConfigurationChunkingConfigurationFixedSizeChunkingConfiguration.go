// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfiguration struct {
	// The maximum number of tokens to include in a chunk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#max_tokens BedrockDataSource#max_tokens}
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// The percentage of overlap between adjacent chunks of a data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#overlap_percentage BedrockDataSource#overlap_percentage}
	OverlapPercentage *float64 `field:"optional" json:"overlapPercentage" yaml:"overlapPercentage"`
}

