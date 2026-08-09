// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog


type GlueCatalogTargetRedshiftCatalog struct {
	// The Amazon Resource Name (ARN) of the catalog resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_catalog#catalog_arn GlueCatalog#catalog_arn}
	CatalogArn *string `field:"optional" json:"catalogArn" yaml:"catalogArn"`
}

