// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource


type AppsyncDataSourceDynamoDbConfigDeltaSyncConfig struct {
	// The number of minutes that an Item is stored in the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_data_source#base_table_ttl AppsyncDataSource#base_table_ttl}
	BaseTableTtl *string `field:"optional" json:"baseTableTtl" yaml:"baseTableTtl"`
	// The Delta Sync table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_data_source#delta_sync_table_name AppsyncDataSource#delta_sync_table_name}
	DeltaSyncTableName *string `field:"optional" json:"deltaSyncTableName" yaml:"deltaSyncTableName"`
	// The number of minutes that a Delta Sync log entry is stored in the Delta Sync table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_data_source#delta_sync_table_ttl AppsyncDataSource#delta_sync_table_ttl}
	DeltaSyncTableTtl *string `field:"optional" json:"deltaSyncTableTtl" yaml:"deltaSyncTableTtl"`
}

