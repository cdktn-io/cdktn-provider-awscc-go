// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationUrls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#seed_url_configuration KendraDataSource#seed_url_configuration}.
	SeedUrlConfiguration *KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfiguration `field:"optional" json:"seedUrlConfiguration" yaml:"seedUrlConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#site_maps_configuration KendraDataSource#site_maps_configuration}.
	SiteMapsConfiguration *KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfiguration `field:"optional" json:"siteMapsConfiguration" yaml:"siteMapsConfiguration"`
}

