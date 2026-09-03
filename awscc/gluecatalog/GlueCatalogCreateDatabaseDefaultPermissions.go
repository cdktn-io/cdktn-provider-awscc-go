// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog


type GlueCatalogCreateDatabaseDefaultPermissions struct {
	// The permissions that are granted to the principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_catalog#permissions GlueCatalog#permissions}
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// The Lake Formation principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_catalog#principal GlueCatalog#principal}
	Principal *GlueCatalogCreateDatabaseDefaultPermissionsPrincipal `field:"optional" json:"principal" yaml:"principal"`
}

