// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionDeadLetterConfig struct {
	// The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_function#target_arn LambdaFunction#target_arn}
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}

