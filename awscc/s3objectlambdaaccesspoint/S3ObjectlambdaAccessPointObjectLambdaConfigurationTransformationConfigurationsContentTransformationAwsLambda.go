// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3objectlambdaaccesspoint


type S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambda struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3objectlambda_access_point#function_arn S3ObjectlambdaAccessPoint#function_arn}.
	FunctionArn *string `field:"optional" json:"functionArn" yaml:"functionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3objectlambda_access_point#function_payload S3ObjectlambdaAccessPoint#function_payload}.
	FunctionPayload *string `field:"optional" json:"functionPayload" yaml:"functionPayload"`
}

