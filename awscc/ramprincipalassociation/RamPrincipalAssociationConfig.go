// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ramprincipalassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RamPrincipalAssociationConfig struct {
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
	// Specifies the principal to associate with the resource share. The possible values are:.
	//
	// - An AWS account ID
	//
	// - An Amazon Resource Name (ARN) of an organization in AWS Organizations
	//
	// - An ARN of an organizational unit (OU) in AWS Organizations
	//
	// - An ARN of an IAM role
	//
	// - An ARN of an IAM user
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/ram_principal_association#principal RamPrincipalAssociation#principal}
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/ram_principal_association#resource_share_arn RamPrincipalAssociation#resource_share_arn}
	ResourceShareArn *string `field:"required" json:"resourceShareArn" yaml:"resourceShareArn"`
}

