// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationtagassociation


type LakeformationTagAssociationLfTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_tag_association#catalog_id LakeformationTagAssociation#catalog_id}.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_tag_association#tag_key LakeformationTagAssociation#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_tag_association#tag_values LakeformationTagAssociation#tag_values}.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
}

