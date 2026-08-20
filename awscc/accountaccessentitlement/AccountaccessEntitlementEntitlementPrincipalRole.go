// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountaccessentitlement


type AccountaccessEntitlementEntitlementPrincipalRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/accountaccess_entitlement#principal AccountaccessEntitlement#principal}.
	Principal *AccountaccessEntitlementEntitlementPrincipalRolePrincipal `field:"required" json:"principal" yaml:"principal"`
	// The ARN of the IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/accountaccess_entitlement#role_arn AccountaccessEntitlement#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

