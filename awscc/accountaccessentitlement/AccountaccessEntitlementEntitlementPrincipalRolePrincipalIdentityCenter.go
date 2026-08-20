// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountaccessentitlement


type AccountaccessEntitlementEntitlementPrincipalRolePrincipalIdentityCenter struct {
	// The ID of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/accountaccess_entitlement#group_id AccountaccessEntitlement#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The ID of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/accountaccess_entitlement#user_id AccountaccessEntitlement#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

