// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessdatasource


type QbusinessDataSourceDocumentEnrichmentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_data_source#inline_configurations QbusinessDataSource#inline_configurations}.
	InlineConfigurations interface{} `field:"optional" json:"inlineConfigurations" yaml:"inlineConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_data_source#post_extraction_hook_configuration QbusinessDataSource#post_extraction_hook_configuration}.
	PostExtractionHookConfiguration *QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfiguration `field:"optional" json:"postExtractionHookConfiguration" yaml:"postExtractionHookConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_data_source#pre_extraction_hook_configuration QbusinessDataSource#pre_extraction_hook_configuration}.
	PreExtractionHookConfiguration *QbusinessDataSourceDocumentEnrichmentConfigurationPreExtractionHookConfiguration `field:"optional" json:"preExtractionHookConfiguration" yaml:"preExtractionHookConfiguration"`
}

