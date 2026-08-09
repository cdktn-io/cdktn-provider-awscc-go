// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonedatasource


type DatazoneDataSourceConfigurationRedshiftRunConfigurationRelationalFilterConfigurations struct {
	// The database name specified in the relational filter configuration for the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_data_source#database_name DatazoneDataSource#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The filter expressions specified in the relational filter configuration for the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_data_source#filter_expressions DatazoneDataSource#filter_expressions}
	FilterExpressions interface{} `field:"optional" json:"filterExpressions" yaml:"filterExpressions"`
	// The schema name specified in the relational filter configuration for the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_data_source#schema_name DatazoneDataSource#schema_name}
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

