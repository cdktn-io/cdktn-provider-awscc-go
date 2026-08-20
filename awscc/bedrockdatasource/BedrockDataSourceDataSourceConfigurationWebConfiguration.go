// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationWebConfiguration struct {
	// Configuration for the web crawler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_data_source#crawler_configuration BedrockDataSource#crawler_configuration}
	CrawlerConfiguration *BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfiguration `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
	// A web source configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_data_source#source_configuration BedrockDataSource#source_configuration}
	SourceConfiguration *BedrockDataSourceDataSourceConfigurationWebConfigurationSourceConfiguration `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

