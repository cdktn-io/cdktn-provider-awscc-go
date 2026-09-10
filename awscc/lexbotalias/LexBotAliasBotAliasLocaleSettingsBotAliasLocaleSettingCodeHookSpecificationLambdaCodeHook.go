// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbotalias


type LexBotAliasBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecificationLambdaCodeHook struct {
	// The version of the request-response that you want Amazon Lex to use to invoke your Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lex_bot_alias#code_hook_interface_version LexBotAlias#code_hook_interface_version}
	CodeHookInterfaceVersion *string `field:"optional" json:"codeHookInterfaceVersion" yaml:"codeHookInterfaceVersion"`
	// The Amazon Resource Name (ARN) of the Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lex_bot_alias#lambda_arn LexBotAlias#lambda_arn}
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
}

