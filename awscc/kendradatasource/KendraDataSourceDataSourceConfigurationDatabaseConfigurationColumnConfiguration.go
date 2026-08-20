// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#change_detecting_columns KendraDataSource#change_detecting_columns}.
	ChangeDetectingColumns *[]*string `field:"optional" json:"changeDetectingColumns" yaml:"changeDetectingColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#document_data_column_name KendraDataSource#document_data_column_name}.
	DocumentDataColumnName *string `field:"optional" json:"documentDataColumnName" yaml:"documentDataColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#document_id_column_name KendraDataSource#document_id_column_name}.
	DocumentIdColumnName *string `field:"optional" json:"documentIdColumnName" yaml:"documentIdColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#document_title_column_name KendraDataSource#document_title_column_name}.
	DocumentTitleColumnName *string `field:"optional" json:"documentTitleColumnName" yaml:"documentTitleColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#field_mappings KendraDataSource#field_mappings}.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}

