// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ramresourceshare


type RamResourceShareResourceShareConfiguration struct {
	// The resource share restricts access to an account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ram_resource_share#exclusive_account_access RamResourceShare#exclusive_account_access}
	ExclusiveAccountAccess interface{} `field:"optional" json:"exclusiveAccountAccess" yaml:"exclusiveAccountAccess"`
	// Specifies whether the consumer account retains access to the resource share after leaving the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ram_resource_share#retain_sharing_on_account_leave_organization RamResourceShare#retain_sharing_on_account_leave_organization}
	RetainSharingOnAccountLeaveOrganization interface{} `field:"optional" json:"retainSharingOnAccountLeaveOrganization" yaml:"retainSharingOnAccountLeaveOrganization"`
}

