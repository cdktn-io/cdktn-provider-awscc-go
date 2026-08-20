// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationprincipalpermissions

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LakeformationPrincipalPermissionsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The permissions granted or revoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lakeformation_principal_permissions#permissions LakeformationPrincipalPermissions#permissions}
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// Indicates the ability to grant permissions (as a subset of permissions granted).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lakeformation_principal_permissions#permissions_with_grant_option LakeformationPrincipalPermissions#permissions_with_grant_option}
	PermissionsWithGrantOption *[]*string `field:"required" json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
	// The principal to be granted a permission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lakeformation_principal_permissions#principal LakeformationPrincipalPermissions#principal}
	Principal *LakeformationPrincipalPermissionsPrincipal `field:"required" json:"principal" yaml:"principal"`
	// The resource to be granted or revoked permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lakeformation_principal_permissions#resource LakeformationPrincipalPermissions#resource}
	Resource *LakeformationPrincipalPermissionsResource `field:"required" json:"resource" yaml:"resource"`
	// The identifier for the GLUDC.
	//
	// By default, the account ID. The GLUDC is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lakeformation_principal_permissions#catalog LakeformationPrincipalPermissions#catalog}
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
}

