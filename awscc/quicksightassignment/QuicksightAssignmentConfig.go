// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightassignment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightAssignmentConfig struct {
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
	// The name of the assignment, also called a rule. The name must be unique within the AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_assignment#assignment_name QuicksightAssignment#assignment_name}
	AssignmentName *string `field:"required" json:"assignmentName" yaml:"assignmentName"`
	// The status of the assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_assignment#assignment_status QuicksightAssignment#assignment_status}
	AssignmentStatus *string `field:"required" json:"assignmentStatus" yaml:"assignmentStatus"`
	// The namespace that contains the assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_assignment#namespace QuicksightAssignment#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The QuickSight users, groups, or both that you want to assign the policy to.
	//
	// This maps identity type to a list of identity names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_assignment#identities QuicksightAssignment#identities}
	Identities interface{} `field:"optional" json:"identities" yaml:"identities"`
	// The ARN for the IAM policy to apply to the QuickSight users and groups specified in this assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/quicksight_assignment#policy_arn QuicksightAssignment#policy_arn}
	PolicyArn *string `field:"optional" json:"policyArn" yaml:"policyArn"`
}

