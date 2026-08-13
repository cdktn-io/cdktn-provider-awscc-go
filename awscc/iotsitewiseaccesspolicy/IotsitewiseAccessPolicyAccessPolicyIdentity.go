// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseaccesspolicy


type IotsitewiseAccessPolicyAccessPolicyIdentity struct {
	// Contains information for an IAM role identity in an access policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_access_policy#iam_role IotsitewiseAccessPolicy#iam_role}
	IamRole *IotsitewiseAccessPolicyAccessPolicyIdentityIamRole `field:"optional" json:"iamRole" yaml:"iamRole"`
	// Contains information for an IAM user identity in an access policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_access_policy#iam_user IotsitewiseAccessPolicy#iam_user}
	IamUser *IotsitewiseAccessPolicyAccessPolicyIdentityIamUser `field:"optional" json:"iamUser" yaml:"iamUser"`
	// Contains information for a user identity in an access policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_access_policy#user IotsitewiseAccessPolicy#user}
	User *IotsitewiseAccessPolicyAccessPolicyIdentityUser `field:"optional" json:"user" yaml:"user"`
}

