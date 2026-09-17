// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationParsingConfiguration struct {
	// Settings for a Bedrock Data Automation used to parse documents for a data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#bedrock_data_automation_configuration BedrockDataSource#bedrock_data_automation_configuration}
	BedrockDataAutomationConfiguration *BedrockDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfiguration `field:"optional" json:"bedrockDataAutomationConfiguration" yaml:"bedrockDataAutomationConfiguration"`
	// Settings for a foundation model used to parse documents for a data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#bedrock_foundation_model_configuration BedrockDataSource#bedrock_foundation_model_configuration}
	BedrockFoundationModelConfiguration *BedrockDataSourceVectorIngestionConfigurationParsingConfigurationBedrockFoundationModelConfiguration `field:"optional" json:"bedrockFoundationModelConfiguration" yaml:"bedrockFoundationModelConfiguration"`
	// The parsing strategy for the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#parsing_strategy BedrockDataSource#parsing_strategy}
	ParsingStrategy *string `field:"optional" json:"parsingStrategy" yaml:"parsingStrategy"`
}

