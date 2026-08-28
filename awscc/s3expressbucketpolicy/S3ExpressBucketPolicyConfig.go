// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3expressbucketpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3ExpressBucketPolicyConfig struct {
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
	// The name of the S3 directory bucket to which the policy applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3express_bucket_policy#bucket S3ExpressBucketPolicy#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// A policy document containing permissions to add to the specified bucket.
	//
	// In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3express_bucket_policy#policy_document S3ExpressBucketPolicy#policy_document}
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
}

