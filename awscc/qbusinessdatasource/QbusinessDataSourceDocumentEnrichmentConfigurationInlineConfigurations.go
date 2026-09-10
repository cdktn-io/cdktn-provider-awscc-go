// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessdatasource


type QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_data_source#condition QbusinessDataSource#condition}.
	Condition *QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_data_source#document_content_operator QbusinessDataSource#document_content_operator}.
	DocumentContentOperator *string `field:"optional" json:"documentContentOperator" yaml:"documentContentOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_data_source#target QbusinessDataSource#target}.
	Target *QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTarget `field:"optional" json:"target" yaml:"target"`
}

