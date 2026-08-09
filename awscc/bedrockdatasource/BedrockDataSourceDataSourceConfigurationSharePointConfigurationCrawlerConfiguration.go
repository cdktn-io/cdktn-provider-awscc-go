// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationSharePointConfigurationCrawlerConfiguration struct {
	// The type of filtering that you want to apply to certain objects or content of the data source.
	//
	// For example, the PATTERN type is regular expression patterns you can apply to filter your content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_source#filter_configuration BedrockDataSource#filter_configuration}
	FilterConfiguration *BedrockDataSourceDataSourceConfigurationSharePointConfigurationCrawlerConfigurationFilterConfiguration `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
}

