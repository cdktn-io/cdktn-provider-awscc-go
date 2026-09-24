// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationContextEnrichmentConfiguration struct {
	// Bedrock Foundation Model configuration to be used for Context Enrichment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#bedrock_foundation_model_configuration BedrockDataSource#bedrock_foundation_model_configuration}
	BedrockFoundationModelConfiguration *BedrockDataSourceVectorIngestionConfigurationContextEnrichmentConfigurationBedrockFoundationModelConfiguration `field:"optional" json:"bedrockFoundationModelConfiguration" yaml:"bedrockFoundationModelConfiguration"`
	// Enrichment type to be used for the vector database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#type BedrockDataSource#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

