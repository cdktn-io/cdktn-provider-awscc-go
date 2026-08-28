// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotTestBotAliasSettingsConversationLogSettingsTextLogSettingsDestinationCloudwatch struct {
	// A string used to identify the groupArn for the Cloudwatch Log Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#cloudwatch_log_group_arn LexBot#cloudwatch_log_group_arn}
	CloudwatchLogGroupArn *string `field:"optional" json:"cloudwatchLogGroupArn" yaml:"cloudwatchLogGroupArn"`
	// A string containing the value for the Log Prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#log_prefix LexBot#log_prefix}
	LogPrefix *string `field:"optional" json:"logPrefix" yaml:"logPrefix"`
}

