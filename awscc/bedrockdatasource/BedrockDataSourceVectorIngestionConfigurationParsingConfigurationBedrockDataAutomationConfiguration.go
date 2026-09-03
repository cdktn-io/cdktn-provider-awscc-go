// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfiguration struct {
	// Determine how will parsed content be stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#parsing_modality BedrockDataSource#parsing_modality}
	ParsingModality *string `field:"optional" json:"parsingModality" yaml:"parsingModality"`
}

