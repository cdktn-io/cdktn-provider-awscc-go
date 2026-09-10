// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamstackuserassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppstreamStackUserAssociationConfig struct {
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
	// The authentication type for the user who is associated with the stack. You must specify USERPOOL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appstream_stack_user_association#authentication_type AppstreamStackUserAssociation#authentication_type}
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The name of the stack that is associated with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appstream_stack_user_association#stack_name AppstreamStackUserAssociation#stack_name}
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// The name of the user who is associated with the stack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appstream_stack_user_association#user_name AppstreamStackUserAssociation#user_name}
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Specifies whether a welcome email is sent to a user after the user is created in the user pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appstream_stack_user_association#send_email_notification AppstreamStackUserAssociation#send_email_notification}
	SendEmailNotification interface{} `field:"optional" json:"sendEmailNotification" yaml:"sendEmailNotification"`
}

