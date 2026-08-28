// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluedatabase


type GlueDatabaseDatabaseInputCreateTableDefaultPermissions struct {
	// The permissions that are granted to the principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_database#permissions GlueDatabase#permissions}
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// The principal who is granted permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_database#principal GlueDatabase#principal}
	Principal *GlueDatabaseDatabaseInputCreateTableDefaultPermissionsPrincipal `field:"optional" json:"principal" yaml:"principal"`
}

