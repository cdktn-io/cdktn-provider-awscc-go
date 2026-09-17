// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpoolregionalconfigurationattachment


type CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigCustomSmsSender struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cognito_user_pool_regional_configuration_attachment#lambda_arn CognitoUserPoolRegionalConfigurationAttachment#lambda_arn}.
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cognito_user_pool_regional_configuration_attachment#lambda_version CognitoUserPoolRegionalConfigurationAttachment#lambda_version}.
	LambdaVersion *string `field:"optional" json:"lambdaVersion" yaml:"lambdaVersion"`
}

