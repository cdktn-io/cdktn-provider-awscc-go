// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluemltransform


type GlueMlTransformInputRecordTablesGlueTables struct {
	// A unique identifier for the AWS Glue Data Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_ml_transform#catalog_id GlueMlTransform#catalog_id}
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the connection to the AWS Glue Data Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_ml_transform#connection_name GlueMlTransform#connection_name}
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// A database name in the AWS Glue Data Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_ml_transform#database_name GlueMlTransform#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// A table name in the AWS Glue Data Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_ml_transform#table_name GlueMlTransform#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

