// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewdataset


type DatabrewDatasetInputDatabaseInputDefinition struct {
	// Database table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_dataset#database_table_name DatabrewDataset#database_table_name}
	DatabaseTableName *string `field:"optional" json:"databaseTableName" yaml:"databaseTableName"`
	// Glue connection name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_dataset#glue_connection_name DatabrewDataset#glue_connection_name}
	GlueConnectionName *string `field:"optional" json:"glueConnectionName" yaml:"glueConnectionName"`
	// Custom SQL to run against the provided AWS Glue connection.
	//
	// This SQL will be used as the input for DataBrew projects and jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_dataset#query_string DatabrewDataset#query_string}
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// Input location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_dataset#temp_directory DatabrewDataset#temp_directory}
	TempDirectory *DatabrewDatasetInputDatabaseInputDefinitionTempDirectory `field:"optional" json:"tempDirectory" yaml:"tempDirectory"`
}

