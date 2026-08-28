// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource


type AppsyncDataSourceRelationalDatabaseConfigRdsHttpEndpointConfig struct {
	// AWS Region for RDS HTTP endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_data_source#aws_region AppsyncDataSource#aws_region}
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The ARN for database credentials stored in AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_data_source#aws_secret_store_arn AppsyncDataSource#aws_secret_store_arn}
	AwsSecretStoreArn *string `field:"optional" json:"awsSecretStoreArn" yaml:"awsSecretStoreArn"`
	// Logical database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_data_source#database_name AppsyncDataSource#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Amazon RDS cluster Amazon Resource Name (ARN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_data_source#db_cluster_identifier AppsyncDataSource#db_cluster_identifier}
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// Logical schema name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_data_source#schema AppsyncDataSource#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

