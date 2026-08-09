// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organizationsaccount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OrganizationsAccountConfig struct {
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
	// The friendly name of the member account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/organizations_account#account_name OrganizationsAccount#account_name}
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// The email address of the owner to assign to the new member account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/organizations_account#email OrganizationsAccount#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// List of parent nodes for the member account.
	//
	// Currently only one parent at a time is supported. Default is root.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/organizations_account#parent_ids OrganizationsAccount#parent_ids}
	ParentIds *[]*string `field:"optional" json:"parentIds" yaml:"parentIds"`
	// The name of an IAM role that AWS Organizations automatically preconfigures in the new member account.
	//
	// Default name is OrganizationAccountAccessRole if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/organizations_account#role_name OrganizationsAccount#role_name}
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// A list of tags that you want to attach to the newly created account.
	//
	// For each tag in the list, you must specify both a tag key and a value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/organizations_account#tags OrganizationsAccount#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

