// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamrolepolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IamRolePolicyConfig struct {
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
	// The name of the policy document.
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iam_role_policy#policy_name IamRolePolicy#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The name of the role to associate the policy with.
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iam_role_policy#role_name IamRolePolicy#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// The policy document.
	//
	// You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.
	//  The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:
	//   +  Any printable ASCII character ranging from the space character (``\u0020``) through the end of the ASCII character range
	//   +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ``\u00FF``)
	//   +  The special characters tab (``\u0009``), line feed (``\u000A``), and carriage return (``\u000D``)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iam_role_policy#policy_document IamRolePolicy#policy_document}
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}

