// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncapi


type AppsyncApiEventConfigAuthProviders struct {
	// Security configuration for your AppSync API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_api#auth_type AppsyncApi#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// Optional authorization configuration for using Amazon Cognito user pools with your API endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_api#cognito_config AppsyncApi#cognito_config}
	CognitoConfig *AppsyncApiEventConfigAuthProvidersCognitoConfig `field:"optional" json:"cognitoConfig" yaml:"cognitoConfig"`
	// A LambdaAuthorizerConfig holds configuration on how to authorize AWS AppSync API access when using the AWS_LAMBDA authorizer mode.
	//
	// Be aware that an AWS AppSync API may have only one Lambda authorizer configured at a time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_api#lambda_authorizer_config AppsyncApi#lambda_authorizer_config}
	LambdaAuthorizerConfig *AppsyncApiEventConfigAuthProvidersLambdaAuthorizerConfig `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// The OpenID Connect configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appsync_api#open_id_connect_config AppsyncApi#open_id_connect_config}
	OpenIdConnectConfig *AppsyncApiEventConfigAuthProvidersOpenIdConnectConfig `field:"optional" json:"openIdConnectConfig" yaml:"openIdConnectConfig"`
}

