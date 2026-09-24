// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationWebConfigurationCrawlerConfigurationCrawlerLimits struct {
	// Maximum number of pages the crawler can crawl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#max_pages BedrockDataSource#max_pages}
	MaxPages *float64 `field:"optional" json:"maxPages" yaml:"maxPages"`
	// Rate of web URLs retrieved per minute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#rate_limit BedrockDataSource#rate_limit}
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

