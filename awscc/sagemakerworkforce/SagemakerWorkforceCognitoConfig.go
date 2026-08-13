// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkforce


type SagemakerWorkforceCognitoConfig struct {
	// The client ID for your Amazon Cognito user pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_workforce#client_id SagemakerWorkforce#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The ID for your Amazon Cognito user pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_workforce#user_pool SagemakerWorkforce#user_pool}
	UserPool *string `field:"optional" json:"userPool" yaml:"userPool"`
}

