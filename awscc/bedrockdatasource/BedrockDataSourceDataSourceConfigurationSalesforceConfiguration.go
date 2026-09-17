// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationSalesforceConfiguration struct {
	// The configuration of filtering the Salesforce content. For example, configuring regular expression patterns to include or exclude certain content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#crawler_configuration BedrockDataSource#crawler_configuration}
	CrawlerConfiguration *BedrockDataSourceDataSourceConfigurationSalesforceConfigurationCrawlerConfiguration `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
	// The endpoint information to connect to your Salesforce data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#source_configuration BedrockDataSource#source_configuration}
	SourceConfiguration *BedrockDataSourceDataSourceConfigurationSalesforceConfigurationSourceConfiguration `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

