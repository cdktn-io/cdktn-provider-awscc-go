// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitoidentitypool


type CognitoIdentityPoolPushSync struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cognito_identity_pool#application_arns CognitoIdentityPool#application_arns}.
	ApplicationArns *[]*string `field:"optional" json:"applicationArns" yaml:"applicationArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cognito_identity_pool#role_arn CognitoIdentityPool#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

