// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewjob


type DatabrewJobDataCatalogOutputs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_job#catalog_id DatabrewJob#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_job#database_name DatabrewJob#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_job#database_options DatabrewJob#database_options}.
	DatabaseOptions *DatabrewJobDataCatalogOutputsDatabaseOptions `field:"optional" json:"databaseOptions" yaml:"databaseOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_job#overwrite DatabrewJob#overwrite}.
	Overwrite interface{} `field:"optional" json:"overwrite" yaml:"overwrite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_job#s3_options DatabrewJob#s3_options}.
	S3Options *DatabrewJobDataCatalogOutputsS3Options `field:"optional" json:"s3Options" yaml:"s3Options"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/databrew_job#table_name DatabrewJob#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

