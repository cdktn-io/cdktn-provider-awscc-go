// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationprincipalpermissions


type LakeformationPrincipalPermissionsResourceDataCellsFilter struct {
	// A database in the GLUDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lakeformation_principal_permissions#database_name LakeformationPrincipalPermissions#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name given by the user to the data filter cell.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lakeformation_principal_permissions#name LakeformationPrincipalPermissions#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the catalog to which the table belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lakeformation_principal_permissions#table_catalog_id LakeformationPrincipalPermissions#table_catalog_id}
	TableCatalogId *string `field:"optional" json:"tableCatalogId" yaml:"tableCatalogId"`
	// The name of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lakeformation_principal_permissions#table_name LakeformationPrincipalPermissions#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

