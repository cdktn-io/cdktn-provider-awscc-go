// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationContextEnrichmentConfigurationBedrockFoundationModelConfiguration struct {
	// Strategy to be used when using Bedrock Foundation Model for Context Enrichment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_source#enrichment_strategy_configuration BedrockDataSource#enrichment_strategy_configuration}
	EnrichmentStrategyConfiguration *BedrockDataSourceVectorIngestionConfigurationContextEnrichmentConfigurationBedrockFoundationModelConfigurationEnrichmentStrategyConfiguration `field:"optional" json:"enrichmentStrategyConfiguration" yaml:"enrichmentStrategyConfiguration"`
	// The model's ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_source#model_arn BedrockDataSource#model_arn}
	ModelArn *string `field:"optional" json:"modelArn" yaml:"modelArn"`
}

