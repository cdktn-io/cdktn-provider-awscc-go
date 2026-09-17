// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceDataSourceConfigurationSalesforceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#chatter_feed_configuration KendraDataSource#chatter_feed_configuration}.
	ChatterFeedConfiguration *KendraDataSourceDataSourceConfigurationSalesforceConfigurationChatterFeedConfiguration `field:"optional" json:"chatterFeedConfiguration" yaml:"chatterFeedConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#crawl_attachments KendraDataSource#crawl_attachments}.
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#exclude_attachment_file_patterns KendraDataSource#exclude_attachment_file_patterns}.
	ExcludeAttachmentFilePatterns *[]*string `field:"optional" json:"excludeAttachmentFilePatterns" yaml:"excludeAttachmentFilePatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#include_attachment_file_patterns KendraDataSource#include_attachment_file_patterns}.
	IncludeAttachmentFilePatterns *[]*string `field:"optional" json:"includeAttachmentFilePatterns" yaml:"includeAttachmentFilePatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#knowledge_article_configuration KendraDataSource#knowledge_article_configuration}.
	KnowledgeArticleConfiguration *KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfiguration `field:"optional" json:"knowledgeArticleConfiguration" yaml:"knowledgeArticleConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#secret_arn KendraDataSource#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#server_url KendraDataSource#server_url}.
	ServerUrl *string `field:"optional" json:"serverUrl" yaml:"serverUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#standard_object_attachment_configuration KendraDataSource#standard_object_attachment_configuration}.
	StandardObjectAttachmentConfiguration *KendraDataSourceDataSourceConfigurationSalesforceConfigurationStandardObjectAttachmentConfiguration `field:"optional" json:"standardObjectAttachmentConfiguration" yaml:"standardObjectAttachmentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#standard_object_configurations KendraDataSource#standard_object_configurations}.
	StandardObjectConfigurations interface{} `field:"optional" json:"standardObjectConfigurations" yaml:"standardObjectConfigurations"`
}

