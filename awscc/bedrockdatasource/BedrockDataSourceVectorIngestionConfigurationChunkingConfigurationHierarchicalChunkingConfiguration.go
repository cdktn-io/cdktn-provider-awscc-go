// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfiguration struct {
	// Token settings for each layer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_source#level_configurations BedrockDataSource#level_configurations}
	LevelConfigurations interface{} `field:"optional" json:"levelConfigurations" yaml:"levelConfigurations"`
	// The number of tokens to repeat across chunks in the same layer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_source#overlap_tokens BedrockDataSource#overlap_tokens}
	OverlapTokens *float64 `field:"optional" json:"overlapTokens" yaml:"overlapTokens"`
}

