// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationConfluenceConfigurationCrawlerConfigurationFilterConfiguration struct {
	// The configuration of specific filters applied to your data source content. You can filter out or include certain content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_source#pattern_object_filter BedrockDataSource#pattern_object_filter}
	PatternObjectFilter *BedrockDataSourceDataSourceConfigurationConfluenceConfigurationCrawlerConfigurationFilterConfigurationPatternObjectFilter `field:"optional" json:"patternObjectFilter" yaml:"patternObjectFilter"`
	// The crawl filter type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_source#type BedrockDataSource#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

