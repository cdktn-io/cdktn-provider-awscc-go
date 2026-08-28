// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaresourcepolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaResourcePolicyConfig struct {
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
	// The policy document you want to add to your LAM resource.
	//
	// This is formatted as a JSON string.
	//  For more information, see [Working with resource-based policies in](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html) in the *Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_resource_policy#policy_document LambdaResourcePolicy#policy_document}
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The Amazon Resource Name (ARN) of the LAM resource you want to add the policy to.
	//
	// For a function, you can use a qualified or an unqualified ARN. The value must be a complete ARN, and the operation does not accept wildcard characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_resource_policy#resource_arn LambdaResourcePolicy#resource_arn}
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

