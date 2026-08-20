// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3vectorsvectorbucketpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3VectorsVectorBucketPolicyConfig struct {
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
	// A policy document containing permissions to add to the specified vector bucket.
	//
	// In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3vectors_vector_bucket_policy#policy S3VectorsVectorBucketPolicy#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// The Amazon Resource Name (ARN) of the vector bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3vectors_vector_bucket_policy#vector_bucket_arn S3VectorsVectorBucketPolicy#vector_bucket_arn}
	VectorBucketArn *string `field:"optional" json:"vectorBucketArn" yaml:"vectorBucketArn"`
	// The name of the vector bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3vectors_vector_bucket_policy#vector_bucket_name S3VectorsVectorBucketPolicy#vector_bucket_name}
	VectorBucketName *string `field:"optional" json:"vectorBucketName" yaml:"vectorBucketName"`
}

