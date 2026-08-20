// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationtagassociation


type LakeformationTagAssociationResourceTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lakeformation_tag_association#catalog_id LakeformationTagAssociation#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lakeformation_tag_association#database_name LakeformationTagAssociation#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lakeformation_tag_association#name LakeformationTagAssociation#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lakeformation_tag_association#table_wildcard LakeformationTagAssociation#table_wildcard}.
	TableWildcard *string `field:"optional" json:"tableWildcard" yaml:"tableWildcard"`
}

