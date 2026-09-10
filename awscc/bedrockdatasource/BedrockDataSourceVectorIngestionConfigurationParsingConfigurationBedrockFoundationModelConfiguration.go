// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationParsingConfigurationBedrockFoundationModelConfiguration struct {
	// The model's ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_data_source#model_arn BedrockDataSource#model_arn}
	ModelArn *string `field:"optional" json:"modelArn" yaml:"modelArn"`
	// Determine how will parsed content be stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_data_source#parsing_modality BedrockDataSource#parsing_modality}
	ParsingModality *string `field:"optional" json:"parsingModality" yaml:"parsingModality"`
	// Instructions for interpreting the contents of a document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_data_source#parsing_prompt BedrockDataSource#parsing_prompt}
	ParsingPrompt *BedrockDataSourceVectorIngestionConfigurationParsingConfigurationBedrockFoundationModelConfigurationParsingPrompt `field:"optional" json:"parsingPrompt" yaml:"parsingPrompt"`
}

