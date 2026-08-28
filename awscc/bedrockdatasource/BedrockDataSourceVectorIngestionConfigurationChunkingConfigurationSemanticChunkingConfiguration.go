// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationSemanticChunkingConfiguration struct {
	// The dissimilarity threshold for splitting chunks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#breakpoint_percentile_threshold BedrockDataSource#breakpoint_percentile_threshold}
	BreakpointPercentileThreshold *float64 `field:"optional" json:"breakpointPercentileThreshold" yaml:"breakpointPercentileThreshold"`
	// The buffer size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#buffer_size BedrockDataSource#buffer_size}
	BufferSize *float64 `field:"optional" json:"bufferSize" yaml:"bufferSize"`
	// The maximum number of tokens that a chunk can contain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#max_tokens BedrockDataSource#max_tokens}
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
}

