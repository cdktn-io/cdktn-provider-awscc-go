// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package detectivememberinvitation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DetectiveMemberInvitationConfig struct {
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
	// The ARN of the graph to which the member account will be invited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/detective_member_invitation#graph_arn DetectiveMemberInvitation#graph_arn}
	GraphArn *string `field:"required" json:"graphArn" yaml:"graphArn"`
	// The root email address for the account to be invited, for validation. Updating this field has no effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/detective_member_invitation#member_email_address DetectiveMemberInvitation#member_email_address}
	MemberEmailAddress *string `field:"required" json:"memberEmailAddress" yaml:"memberEmailAddress"`
	// The AWS account ID to be invited to join the graph as a member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/detective_member_invitation#member_id DetectiveMemberInvitation#member_id}
	MemberId *string `field:"required" json:"memberId" yaml:"memberId"`
	// When set to true, invitation emails are not sent to the member accounts.
	//
	// Member accounts must still accept the invitation before they are added to the behavior graph. Updating this field has no effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/detective_member_invitation#disable_email_notification DetectiveMemberInvitation#disable_email_notification}
	DisableEmailNotification interface{} `field:"optional" json:"disableEmailNotification" yaml:"disableEmailNotification"`
	// A message to be included in the email invitation sent to the invited account.
	//
	// Updating this field has no effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/detective_member_invitation#message DetectiveMemberInvitation#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
}

