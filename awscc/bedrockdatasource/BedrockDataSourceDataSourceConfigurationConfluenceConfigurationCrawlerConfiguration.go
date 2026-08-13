// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationConfluenceConfigurationCrawlerConfiguration struct {
	// The type of filtering that you want to apply to certain objects or content of the data source.
	//
	// For example, the PATTERN type is regular expression patterns you can apply to filter your content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_data_source#filter_configuration BedrockDataSource#filter_configuration}
	FilterConfiguration *BedrockDataSourceDataSourceConfigurationConfluenceConfigurationCrawlerConfigurationFilterConfiguration `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
}

