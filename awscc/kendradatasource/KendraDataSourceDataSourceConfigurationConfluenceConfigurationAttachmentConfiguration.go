// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceDataSourceConfigurationConfluenceConfigurationAttachmentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#attachment_field_mappings KendraDataSource#attachment_field_mappings}.
	AttachmentFieldMappings interface{} `field:"optional" json:"attachmentFieldMappings" yaml:"attachmentFieldMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#crawl_attachments KendraDataSource#crawl_attachments}.
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
}

