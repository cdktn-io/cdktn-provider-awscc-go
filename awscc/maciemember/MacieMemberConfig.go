// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package maciemember

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MacieMemberConfig struct {
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
	// The AWS account ID for the account to associate with the Amazon Macie administrator account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/macie_member#account_id MacieMember#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The email address for the account to associate with the Amazon Macie administrator account.
	//
	// Required by the Amazon Macie CreateMember API at creation time; it is write-only because the service does not return it (it is null when the account is associated through AWS Organizations).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/macie_member#email MacieMember#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// The tags to associate with the member account in Amazon Macie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/macie_member#tags MacieMember#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

