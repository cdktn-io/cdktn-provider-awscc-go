// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationtagassociation


type LakeformationTagAssociationResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lakeformation_tag_association#catalog LakeformationTagAssociation#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lakeformation_tag_association#database LakeformationTagAssociation#database}.
	Database *LakeformationTagAssociationResourceDatabase `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lakeformation_tag_association#table LakeformationTagAssociation#table}.
	Table *LakeformationTagAssociationResourceTable `field:"optional" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lakeformation_tag_association#table_with_columns LakeformationTagAssociation#table_with_columns}.
	TableWithColumns *LakeformationTagAssociationResourceTableWithColumns `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}

