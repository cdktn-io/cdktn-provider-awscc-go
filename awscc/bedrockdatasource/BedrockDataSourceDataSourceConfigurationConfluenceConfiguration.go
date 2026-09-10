// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationConfluenceConfiguration struct {
	// The configuration of the Confluence content. For example, configuring specific types of Confluence content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_data_source#crawler_configuration BedrockDataSource#crawler_configuration}
	CrawlerConfiguration *BedrockDataSourceDataSourceConfigurationConfluenceConfigurationCrawlerConfiguration `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
	// The endpoint information to connect to your Confluence data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_data_source#source_configuration BedrockDataSource#source_configuration}
	SourceConfiguration *BedrockDataSourceDataSourceConfigurationConfluenceConfigurationSourceConfiguration `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

