// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfiguration struct {
	// A location for storing content from data sources temporarily as it is processed by custom components in the ingestion pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#intermediate_storage BedrockDataSource#intermediate_storage}
	IntermediateStorage *BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationIntermediateStorage `field:"optional" json:"intermediateStorage" yaml:"intermediateStorage"`
	// A list of Lambda functions that process documents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#transformations BedrockDataSource#transformations}
	Transformations interface{} `field:"optional" json:"transformations" yaml:"transformations"`
}

