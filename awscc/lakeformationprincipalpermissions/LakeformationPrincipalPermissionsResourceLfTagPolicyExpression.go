// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationprincipalpermissions


type LakeformationPrincipalPermissionsResourceLfTagPolicyExpression struct {
	// The key-name for the LF-tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lakeformation_principal_permissions#tag_key LakeformationPrincipalPermissions#tag_key}
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// A list of possible values of the corresponding ``TagKey`` of an LF-tag key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lakeformation_principal_permissions#tag_values LakeformationPrincipalPermissions#tag_values}
	TagValues *[]*string `field:"optional" json:"tagValues" yaml:"tagValues"`
}

