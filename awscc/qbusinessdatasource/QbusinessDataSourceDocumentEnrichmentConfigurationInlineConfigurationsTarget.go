// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessdatasource


type QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/qbusiness_data_source#attribute_value_operator QbusinessDataSource#attribute_value_operator}.
	AttributeValueOperator *string `field:"optional" json:"attributeValueOperator" yaml:"attributeValueOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/qbusiness_data_source#key QbusinessDataSource#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/qbusiness_data_source#value QbusinessDataSource#value}.
	Value *QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetValue `field:"optional" json:"value" yaml:"value"`
}

