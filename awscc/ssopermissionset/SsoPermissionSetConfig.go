// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssopermissionset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsoPermissionSetConfig struct {
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
	// The sso instance arn that the permission set is owned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sso_permission_set#instance_arn SsoPermissionSet#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name you want to assign to this permission set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sso_permission_set#name SsoPermissionSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sso_permission_set#customer_managed_policy_references SsoPermissionSet#customer_managed_policy_references}.
	CustomerManagedPolicyReferences interface{} `field:"optional" json:"customerManagedPolicyReferences" yaml:"customerManagedPolicyReferences"`
	// The permission set description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sso_permission_set#description SsoPermissionSet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The inline policy to put in permission set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sso_permission_set#inline_policy SsoPermissionSet#inline_policy}
	InlinePolicy *string `field:"optional" json:"inlinePolicy" yaml:"inlinePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sso_permission_set#managed_policies SsoPermissionSet#managed_policies}.
	ManagedPolicies *[]*string `field:"optional" json:"managedPolicies" yaml:"managedPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sso_permission_set#permissions_boundary SsoPermissionSet#permissions_boundary}.
	PermissionsBoundary *SsoPermissionSetPermissionsBoundary `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// The relay state URL that redirect links to any service in the AWS Management Console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sso_permission_set#relay_state_type SsoPermissionSet#relay_state_type}
	RelayStateType *string `field:"optional" json:"relayStateType" yaml:"relayStateType"`
	// The length of time that a user can be signed in to an AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sso_permission_set#session_duration SsoPermissionSet#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sso_permission_set#tags SsoPermissionSet#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

