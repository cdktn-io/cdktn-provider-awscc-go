// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog


type GlueCatalogCatalogPropertiesDataLakeAccessProperties struct {
	// Allows third-party engines to access data in Amazon S3 locations that are registered with Lake Formation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_catalog#allow_full_table_external_data_access GlueCatalog#allow_full_table_external_data_access}
	AllowFullTableExternalDataAccess *string `field:"optional" json:"allowFullTableExternalDataAccess" yaml:"allowFullTableExternalDataAccess"`
	// Specifies a federated catalog type for the native catalog resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_catalog#catalog_type GlueCatalog#catalog_type}
	CatalogType *string `field:"optional" json:"catalogType" yaml:"catalogType"`
	// Turns on or off data lake access for Apache Spark applications that access Amazon Redshift databases in the Data Catalog from any non-Redshift engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_catalog#data_lake_access GlueCatalog#data_lake_access}
	DataLakeAccess interface{} `field:"optional" json:"dataLakeAccess" yaml:"dataLakeAccess"`
	// A role that will be assumed by Glue for transferring data into/out of the staging bucket during a query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_catalog#data_transfer_role GlueCatalog#data_transfer_role}
	DataTransferRole *string `field:"optional" json:"dataTransferRole" yaml:"dataTransferRole"`
	// An encryption key that will be used for the staging bucket that will be created along with the catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_catalog#kms_key GlueCatalog#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

