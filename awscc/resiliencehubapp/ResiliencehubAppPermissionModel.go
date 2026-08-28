// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubapp


type ResiliencehubAppPermissionModel struct {
	// Defines a list of role Amazon Resource Names (ARNs) to be used in other accounts.
	//
	// These ARNs are used for querying purposes while importing resources and assessing your application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resiliencehub_app#cross_account_role_arns ResiliencehubApp#cross_account_role_arns}
	CrossAccountRoleArns *[]*string `field:"optional" json:"crossAccountRoleArns" yaml:"crossAccountRoleArns"`
	// Existing AWS IAM role name in the primary AWS account that will be assumed by AWS Resilience Hub Service Principle to obtain a read-only access to your application resources while running an assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resiliencehub_app#invoker_role_name ResiliencehubApp#invoker_role_name}
	InvokerRoleName *string `field:"optional" json:"invokerRoleName" yaml:"invokerRoleName"`
	// Defines how AWS Resilience Hub scans your resources.
	//
	// It can scan for the resources by using a pre-existing role in your AWS account, or by using the credentials of the current IAM user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resiliencehub_app#type ResiliencehubApp#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

