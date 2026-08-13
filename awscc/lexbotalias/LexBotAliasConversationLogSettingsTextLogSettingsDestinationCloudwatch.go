// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbotalias


type LexBotAliasConversationLogSettingsTextLogSettingsDestinationCloudwatch struct {
	// A string used to identify the groupArn for the Cloudwatch Log Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot_alias#cloudwatch_log_group_arn LexBotAlias#cloudwatch_log_group_arn}
	CloudwatchLogGroupArn *string `field:"optional" json:"cloudwatchLogGroupArn" yaml:"cloudwatchLogGroupArn"`
	// A string containing the value for the Log Prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot_alias#log_prefix LexBotAlias#log_prefix}
	LogPrefix *string `field:"optional" json:"logPrefix" yaml:"logPrefix"`
}

