// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceDataSourceConfigurationDatabaseConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#acl_configuration KendraDataSource#acl_configuration}.
	AclConfiguration *KendraDataSourceDataSourceConfigurationDatabaseConfigurationAclConfiguration `field:"optional" json:"aclConfiguration" yaml:"aclConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#column_configuration KendraDataSource#column_configuration}.
	ColumnConfiguration *KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfiguration `field:"optional" json:"columnConfiguration" yaml:"columnConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#connection_configuration KendraDataSource#connection_configuration}.
	ConnectionConfiguration *KendraDataSourceDataSourceConfigurationDatabaseConfigurationConnectionConfiguration `field:"optional" json:"connectionConfiguration" yaml:"connectionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#database_engine_type KendraDataSource#database_engine_type}.
	DatabaseEngineType *string `field:"optional" json:"databaseEngineType" yaml:"databaseEngineType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#sql_configuration KendraDataSource#sql_configuration}.
	SqlConfiguration *KendraDataSourceDataSourceConfigurationDatabaseConfigurationSqlConfiguration `field:"optional" json:"sqlConfiguration" yaml:"sqlConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kendra_data_source#vpc_configuration KendraDataSource#vpc_configuration}.
	VpcConfiguration *KendraDataSourceDataSourceConfigurationDatabaseConfigurationVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

