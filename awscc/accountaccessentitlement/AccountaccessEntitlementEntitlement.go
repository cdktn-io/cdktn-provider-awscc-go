// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountaccessentitlement


type AccountaccessEntitlementEntitlement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/accountaccess_entitlement#principal_role AccountaccessEntitlement#principal_role}.
	PrincipalRole *AccountaccessEntitlementEntitlementPrincipalRole `field:"required" json:"principalRole" yaml:"principalRole"`
}

