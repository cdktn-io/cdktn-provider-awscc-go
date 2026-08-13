// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessdatasource


type QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsTargetValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_data_source#date_value QbusinessDataSource#date_value}.
	DateValue *string `field:"optional" json:"dateValue" yaml:"dateValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_data_source#long_value QbusinessDataSource#long_value}.
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_data_source#string_list_value QbusinessDataSource#string_list_value}.
	StringListValue *[]*string `field:"optional" json:"stringListValue" yaml:"stringListValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_data_source#string_value QbusinessDataSource#string_value}.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

