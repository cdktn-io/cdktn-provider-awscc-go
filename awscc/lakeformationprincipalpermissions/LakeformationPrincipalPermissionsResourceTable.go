// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationprincipalpermissions


type LakeformationPrincipalPermissionsResourceTable struct {
	// The identifier for the Data Catalog. By default, it is the account ID of the caller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lakeformation_principal_permissions#catalog_id LakeformationPrincipalPermissions#catalog_id}
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the database for the table.
	//
	// Unique to a Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lakeformation_principal_permissions#database_name LakeformationPrincipalPermissions#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lakeformation_principal_permissions#name LakeformationPrincipalPermissions#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A wildcard object representing every table under a database.  At least one of ``TableResource$Name`` or ``TableResource$TableWildcard`` is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lakeformation_principal_permissions#table_wildcard LakeformationPrincipalPermissions#table_wildcard}
	TableWildcard *string `field:"optional" json:"tableWildcard" yaml:"tableWildcard"`
}

