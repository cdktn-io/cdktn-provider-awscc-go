// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssoapplicationassignment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsoApplicationAssignmentConfig struct {
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
	// The ARN of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sso_application_assignment#application_arn SsoApplicationAssignment#application_arn}
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// An identifier for an object in IAM Identity Center, such as a user or group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sso_application_assignment#principal_id SsoApplicationAssignment#principal_id}
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The entity type for which the assignment will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sso_application_assignment#principal_type SsoApplicationAssignment#principal_type}
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
}

