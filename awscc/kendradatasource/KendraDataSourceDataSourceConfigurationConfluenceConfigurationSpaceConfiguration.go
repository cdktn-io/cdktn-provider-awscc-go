// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#crawl_archived_spaces KendraDataSource#crawl_archived_spaces}.
	CrawlArchivedSpaces interface{} `field:"optional" json:"crawlArchivedSpaces" yaml:"crawlArchivedSpaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#crawl_personal_spaces KendraDataSource#crawl_personal_spaces}.
	CrawlPersonalSpaces interface{} `field:"optional" json:"crawlPersonalSpaces" yaml:"crawlPersonalSpaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#exclude_spaces KendraDataSource#exclude_spaces}.
	ExcludeSpaces *[]*string `field:"optional" json:"excludeSpaces" yaml:"excludeSpaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#include_spaces KendraDataSource#include_spaces}.
	IncludeSpaces *[]*string `field:"optional" json:"includeSpaces" yaml:"includeSpaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_data_source#space_field_mappings KendraDataSource#space_field_mappings}.
	SpaceFieldMappings interface{} `field:"optional" json:"spaceFieldMappings" yaml:"spaceFieldMappings"`
}

