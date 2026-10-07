// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdawebfunctionrevision


type LambdaWebFunctionRevisionBuildConfigCodeConfigS3Object struct {
	// The S3 bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/lambda_web_function_revision#bucket LambdaWebFunctionRevision#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The S3 object key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/lambda_web_function_revision#key LambdaWebFunctionRevision#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The S3 object version ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/lambda_web_function_revision#version_id LambdaWebFunctionRevision#version_id}
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

