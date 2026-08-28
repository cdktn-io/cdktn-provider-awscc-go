// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfigurationLevelConfigurations struct {
	// The maximum number of tokens that a chunk can contain in this layer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#max_tokens BedrockDataSource#max_tokens}
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
}

