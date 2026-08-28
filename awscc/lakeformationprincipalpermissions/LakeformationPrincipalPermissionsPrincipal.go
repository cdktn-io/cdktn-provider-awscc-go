// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationprincipalpermissions


type LakeformationPrincipalPermissionsPrincipal struct {
	// An identifier for the LFlong principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_principal_permissions#data_lake_principal_identifier LakeformationPrincipalPermissions#data_lake_principal_identifier}
	DataLakePrincipalIdentifier *string `field:"optional" json:"dataLakePrincipalIdentifier" yaml:"dataLakePrincipalIdentifier"`
}

