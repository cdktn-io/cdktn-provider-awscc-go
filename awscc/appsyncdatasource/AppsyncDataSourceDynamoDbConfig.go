// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource


type AppsyncDataSourceDynamoDbConfig struct {
	// The AWS Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appsync_data_source#aws_region AppsyncDataSource#aws_region}
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The DeltaSyncConfig for a versioned datasource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appsync_data_source#delta_sync_config AppsyncDataSource#delta_sync_config}
	DeltaSyncConfig *AppsyncDataSourceDynamoDbConfigDeltaSyncConfig `field:"optional" json:"deltaSyncConfig" yaml:"deltaSyncConfig"`
	// The table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appsync_data_source#table_name AppsyncDataSource#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Set to TRUE to use AWS Identity and Access Management with this data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appsync_data_source#use_caller_credentials AppsyncDataSource#use_caller_credentials}
	UseCallerCredentials interface{} `field:"optional" json:"useCallerCredentials" yaml:"useCallerCredentials"`
	// Set to TRUE to use Conflict Detection and Resolution with this data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appsync_data_source#versioned AppsyncDataSource#versioned}
	Versioned interface{} `field:"optional" json:"versioned" yaml:"versioned"`
}

