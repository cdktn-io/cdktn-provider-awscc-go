// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluedatabase


type GlueDatabaseDatabaseInputTargetDatabase struct {
	// The ID of the Data Catalog in which the database resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_database#catalog_id GlueDatabase#catalog_id}
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the catalog database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_database#database_name GlueDatabase#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Region of the target database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_database#region GlueDatabase#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

