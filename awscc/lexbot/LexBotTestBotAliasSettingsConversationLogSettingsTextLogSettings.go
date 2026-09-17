// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotTestBotAliasSettingsConversationLogSettingsTextLogSettings struct {
	// Defines the Amazon CloudWatch Logs destination log group for conversation text logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#destination LexBot#destination}
	Destination *LexBotTestBotAliasSettingsConversationLogSettingsTextLogSettingsDestination `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#enabled LexBot#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

