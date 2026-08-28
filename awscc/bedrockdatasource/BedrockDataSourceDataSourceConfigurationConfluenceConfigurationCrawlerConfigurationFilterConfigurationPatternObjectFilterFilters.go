// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationConfluenceConfigurationCrawlerConfigurationFilterConfigurationPatternObjectFilterFilters struct {
	// A set of regular expression filter patterns for a type of object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#exclusion_filters BedrockDataSource#exclusion_filters}
	ExclusionFilters *[]*string `field:"optional" json:"exclusionFilters" yaml:"exclusionFilters"`
	// A set of regular expression filter patterns for a type of object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#inclusion_filters BedrockDataSource#inclusion_filters}
	InclusionFilters *[]*string `field:"optional" json:"inclusionFilters" yaml:"inclusionFilters"`
	// The supported object type or content type of the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_data_source#object_type BedrockDataSource#object_type}
	ObjectType *string `field:"optional" json:"objectType" yaml:"objectType"`
}

