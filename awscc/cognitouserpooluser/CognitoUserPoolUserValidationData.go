// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpooluser


type CognitoUserPoolUserValidationData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cognito_user_pool_user#name CognitoUserPoolUser#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cognito_user_pool_user#value CognitoUserPoolUser#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

