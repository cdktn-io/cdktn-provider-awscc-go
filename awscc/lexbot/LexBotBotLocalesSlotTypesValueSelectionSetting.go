// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesSlotTypesValueSelectionSetting struct {
	// Provides settings that enable advanced recognition settings for slot values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#advanced_recognition_setting LexBot#advanced_recognition_setting}
	AdvancedRecognitionSetting *LexBotBotLocalesSlotTypesValueSelectionSettingAdvancedRecognitionSetting `field:"optional" json:"advancedRecognitionSetting" yaml:"advancedRecognitionSetting"`
	// A regular expression used to validate the value of a slot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#regex_filter LexBot#regex_filter}
	RegexFilter *LexBotBotLocalesSlotTypesValueSelectionSettingRegexFilter `field:"optional" json:"regexFilter" yaml:"regexFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#resolution_strategy LexBot#resolution_strategy}.
	ResolutionStrategy *string `field:"optional" json:"resolutionStrategy" yaml:"resolutionStrategy"`
}

