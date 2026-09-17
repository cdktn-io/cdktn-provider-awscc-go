// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationprincipalpermissions


type LakeformationPrincipalPermissionsResourceDatabase struct {
	// The identifier for the Data Catalog. By default, it is the account ID of the caller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_principal_permissions#catalog_id LakeformationPrincipalPermissions#catalog_id}
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the database resource. Unique to the Data Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_principal_permissions#name LakeformationPrincipalPermissions#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

