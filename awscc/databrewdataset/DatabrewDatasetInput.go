// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewdataset


type DatabrewDatasetInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/databrew_dataset#database_input_definition DatabrewDataset#database_input_definition}.
	DatabaseInputDefinition *DatabrewDatasetInputDatabaseInputDefinition `field:"optional" json:"databaseInputDefinition" yaml:"databaseInputDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/databrew_dataset#data_catalog_input_definition DatabrewDataset#data_catalog_input_definition}.
	DataCatalogInputDefinition *DatabrewDatasetInputDataCatalogInputDefinition `field:"optional" json:"dataCatalogInputDefinition" yaml:"dataCatalogInputDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/databrew_dataset#metadata DatabrewDataset#metadata}.
	Metadata *DatabrewDatasetInputMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Input location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/databrew_dataset#s3_input_definition DatabrewDataset#s3_input_definition}
	S3InputDefinition *DatabrewDatasetInputS3InputDefinition `field:"optional" json:"s3InputDefinition" yaml:"s3InputDefinition"`
}

