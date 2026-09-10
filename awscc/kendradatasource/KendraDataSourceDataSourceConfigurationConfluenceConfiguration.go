// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceDataSourceConfigurationConfluenceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#attachment_configuration KendraDataSource#attachment_configuration}.
	AttachmentConfiguration *KendraDataSourceDataSourceConfigurationConfluenceConfigurationAttachmentConfiguration `field:"optional" json:"attachmentConfiguration" yaml:"attachmentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#blog_configuration KendraDataSource#blog_configuration}.
	BlogConfiguration *KendraDataSourceDataSourceConfigurationConfluenceConfigurationBlogConfiguration `field:"optional" json:"blogConfiguration" yaml:"blogConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#exclusion_patterns KendraDataSource#exclusion_patterns}.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#inclusion_patterns KendraDataSource#inclusion_patterns}.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#page_configuration KendraDataSource#page_configuration}.
	PageConfiguration *KendraDataSourceDataSourceConfigurationConfluenceConfigurationPageConfiguration `field:"optional" json:"pageConfiguration" yaml:"pageConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#secret_arn KendraDataSource#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#server_url KendraDataSource#server_url}.
	ServerUrl *string `field:"optional" json:"serverUrl" yaml:"serverUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#space_configuration KendraDataSource#space_configuration}.
	SpaceConfiguration *KendraDataSourceDataSourceConfigurationConfluenceConfigurationSpaceConfiguration `field:"optional" json:"spaceConfiguration" yaml:"spaceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#version KendraDataSource#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_data_source#vpc_configuration KendraDataSource#vpc_configuration}.
	VpcConfiguration *KendraDataSourceDataSourceConfigurationConfluenceConfigurationVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

