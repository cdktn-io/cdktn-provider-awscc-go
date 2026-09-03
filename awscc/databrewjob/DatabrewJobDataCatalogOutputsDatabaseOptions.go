// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewjob


type DatabrewJobDataCatalogOutputsDatabaseOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#table_name DatabrewJob#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// S3 Output location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#temp_directory DatabrewJob#temp_directory}
	TempDirectory *DatabrewJobDataCatalogOutputsDatabaseOptionsTempDirectory `field:"optional" json:"tempDirectory" yaml:"tempDirectory"`
}

