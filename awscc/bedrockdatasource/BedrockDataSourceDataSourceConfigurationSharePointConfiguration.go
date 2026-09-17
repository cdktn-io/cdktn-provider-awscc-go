// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationSharePointConfiguration struct {
	// The configuration of the SharePoint content. For example, configuring specific types of SharePoint content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#crawler_configuration BedrockDataSource#crawler_configuration}
	CrawlerConfiguration *BedrockDataSourceDataSourceConfigurationSharePointConfigurationCrawlerConfiguration `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
	// The endpoint information to connect to your SharePoint data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#source_configuration BedrockDataSource#source_configuration}
	SourceConfiguration *BedrockDataSourceDataSourceConfigurationSharePointConfigurationSourceConfiguration `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

