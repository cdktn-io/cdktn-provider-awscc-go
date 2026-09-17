// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationtagassociation


type LakeformationTagAssociationResourceTableWithColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_tag_association#catalog_id LakeformationTagAssociation#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_tag_association#column_names LakeformationTagAssociation#column_names}.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_tag_association#database_name LakeformationTagAssociation#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lakeformation_tag_association#name LakeformationTagAssociation#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

