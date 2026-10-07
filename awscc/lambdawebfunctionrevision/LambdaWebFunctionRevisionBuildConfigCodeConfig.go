// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdawebfunctionrevision


type LambdaWebFunctionRevisionBuildConfigCodeConfig struct {
	// The Amazon S3 location of the deployment artifact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/lambda_web_function_revision#s3_object LambdaWebFunctionRevision#s3_object}
	S3Object *LambdaWebFunctionRevisionBuildConfigCodeConfigS3Object `field:"required" json:"s3Object" yaml:"s3Object"`
}

