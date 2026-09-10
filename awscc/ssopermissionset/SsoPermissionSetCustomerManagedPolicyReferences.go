// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssopermissionset


type SsoPermissionSetCustomerManagedPolicyReferences struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sso_permission_set#name SsoPermissionSet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sso_permission_set#path SsoPermissionSet#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

