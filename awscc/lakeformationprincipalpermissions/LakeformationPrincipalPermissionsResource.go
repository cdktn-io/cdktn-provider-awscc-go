// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationprincipalpermissions


type LakeformationPrincipalPermissionsResource struct {
	// The identifier for the Data Catalog.
	//
	// By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your LFlong environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_principal_permissions#catalog LakeformationPrincipalPermissions#catalog}
	Catalog *LakeformationPrincipalPermissionsResourceCatalog `field:"optional" json:"catalog" yaml:"catalog"`
	// The database for the resource.
	//
	// Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database permissions to a principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_principal_permissions#database LakeformationPrincipalPermissions#database}
	Database *LakeformationPrincipalPermissionsResourceDatabase `field:"optional" json:"database" yaml:"database"`
	// A data cell filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_principal_permissions#data_cells_filter LakeformationPrincipalPermissions#data_cells_filter}
	DataCellsFilter *LakeformationPrincipalPermissionsResourceDataCellsFilter `field:"optional" json:"dataCellsFilter" yaml:"dataCellsFilter"`
	// The location of an Amazon S3 path where permissions are granted or revoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_principal_permissions#data_location LakeformationPrincipalPermissions#data_location}
	DataLocation *LakeformationPrincipalPermissionsResourceDataLocation `field:"optional" json:"dataLocation" yaml:"dataLocation"`
	// The LF-tag key and values attached to a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_principal_permissions#lf_tag LakeformationPrincipalPermissions#lf_tag}
	LfTag *LakeformationPrincipalPermissionsResourceLfTag `field:"optional" json:"lfTag" yaml:"lfTag"`
	// A list of LF-tag conditions that define a resource's LF-tag policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_principal_permissions#lf_tag_policy LakeformationPrincipalPermissions#lf_tag_policy}
	LfTagPolicy *LakeformationPrincipalPermissionsResourceLfTagPolicy `field:"optional" json:"lfTagPolicy" yaml:"lfTagPolicy"`
	// The table for the resource.
	//
	// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_principal_permissions#table LakeformationPrincipalPermissions#table}
	Table *LakeformationPrincipalPermissionsResourceTable `field:"optional" json:"table" yaml:"table"`
	// The table with columns for the resource.
	//
	// A principal with permissions to this resource can select metadata from the columns of a table in the Data Catalog and the underlying data in Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_principal_permissions#table_with_columns LakeformationPrincipalPermissions#table_with_columns}
	TableWithColumns *LakeformationPrincipalPermissionsResourceTableWithColumns `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}

