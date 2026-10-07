// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkteam


type SagemakerWorkteamMemberDefinitionsCognitoMemberDefinition struct {
	// An identifier for an application client. You must create the app client ID using Amazon Cognito.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/sagemaker_workteam#cognito_client_id SagemakerWorkteam#cognito_client_id}
	CognitoClientId *string `field:"optional" json:"cognitoClientId" yaml:"cognitoClientId"`
	// An identifier for a user group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/sagemaker_workteam#cognito_user_group SagemakerWorkteam#cognito_user_group}
	CognitoUserGroup *string `field:"optional" json:"cognitoUserGroup" yaml:"cognitoUserGroup"`
	// An identifier for a user pool.
	//
	// The user pool must be in the same region as the service that you are calling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/sagemaker_workteam#cognito_user_pool SagemakerWorkteam#cognito_user_pool}
	CognitoUserPool *string `field:"optional" json:"cognitoUserPool" yaml:"cognitoUserPool"`
}

