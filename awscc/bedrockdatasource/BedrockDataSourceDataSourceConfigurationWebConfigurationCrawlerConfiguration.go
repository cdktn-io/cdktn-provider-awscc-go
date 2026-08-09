// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfiguration struct {
	// Limit settings for the web crawler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_source#crawler_limits BedrockDataSource#crawler_limits}
	CrawlerLimits *BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationCrawlerLimits `field:"optional" json:"crawlerLimits" yaml:"crawlerLimits"`
	// A set of regular expression filter patterns for a type of object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_source#exclusion_filters BedrockDataSource#exclusion_filters}
	ExclusionFilters *[]*string `field:"optional" json:"exclusionFilters" yaml:"exclusionFilters"`
	// A set of regular expression filter patterns for a type of object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_source#inclusion_filters BedrockDataSource#inclusion_filters}
	InclusionFilters *[]*string `field:"optional" json:"inclusionFilters" yaml:"inclusionFilters"`
	// The scope that a web crawl job will be restricted to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_source#scope BedrockDataSource#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The suffix that will be included in the user agent header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_source#user_agent BedrockDataSource#user_agent}
	UserAgent *string `field:"optional" json:"userAgent" yaml:"userAgent"`
}

