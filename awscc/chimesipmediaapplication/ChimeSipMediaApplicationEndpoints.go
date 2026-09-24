// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimesipmediaapplication


type ChimeSipMediaApplicationEndpoints struct {
	// Valid Amazon Resource Name (ARN) of the Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_sip_media_application#lambda_arn ChimeSipMediaApplication#lambda_arn}
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
}

