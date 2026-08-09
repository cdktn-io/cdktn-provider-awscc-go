// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewdataset


type DatabrewDatasetInputDataCatalogInputDefinition struct {
	// Catalog id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_dataset#catalog_id DatabrewDataset#catalog_id}
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_dataset#database_name DatabrewDataset#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_dataset#table_name DatabrewDataset#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Input location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_dataset#temp_directory DatabrewDataset#temp_directory}
	TempDirectory *DatabrewDatasetInputDataCatalogInputDefinitionTempDirectory `field:"optional" json:"tempDirectory" yaml:"tempDirectory"`
}

