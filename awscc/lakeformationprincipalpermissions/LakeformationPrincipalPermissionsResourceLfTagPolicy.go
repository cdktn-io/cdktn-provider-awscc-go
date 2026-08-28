// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationprincipalpermissions


type LakeformationPrincipalPermissionsResourceLfTagPolicy struct {
	// The identifier for the GLUDC.
	//
	// The GLUDC is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your LFlong environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_principal_permissions#catalog_id LakeformationPrincipalPermissions#catalog_id}
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// A list of LF-tag conditions that apply to the resource's LF-tag policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_principal_permissions#expression LakeformationPrincipalPermissions#expression}
	Expression interface{} `field:"optional" json:"expression" yaml:"expression"`
	// The resource type for which the LF-tag policy applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_principal_permissions#resource_type LakeformationPrincipalPermissions#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

