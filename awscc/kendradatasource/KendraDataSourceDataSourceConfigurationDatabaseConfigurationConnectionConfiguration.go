// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceDataSourceConfigurationDatabaseConfigurationConnectionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#database_host KendraDataSource#database_host}.
	DatabaseHost *string `field:"optional" json:"databaseHost" yaml:"databaseHost"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#database_name KendraDataSource#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#database_port KendraDataSource#database_port}.
	DatabasePort *float64 `field:"optional" json:"databasePort" yaml:"databasePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#secret_arn KendraDataSource#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_data_source#table_name KendraDataSource#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

