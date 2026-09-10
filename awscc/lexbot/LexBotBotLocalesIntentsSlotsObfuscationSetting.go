// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsSlotsObfuscationSetting struct {
	// Value that determines whether Amazon Lex obscures slot values in conversation logs. The default is to obscure the values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lex_bot#obfuscation_setting_type LexBot#obfuscation_setting_type}
	ObfuscationSettingType *string `field:"optional" json:"obfuscationSettingType" yaml:"obfuscationSettingType"`
}

