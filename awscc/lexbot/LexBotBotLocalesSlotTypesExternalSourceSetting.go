// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesSlotTypesExternalSourceSetting struct {
	// Settings required for a slot type based on a grammar that you provide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot#grammar_slot_type_setting LexBot#grammar_slot_type_setting}
	GrammarSlotTypeSetting *LexBotBotLocalesSlotTypesExternalSourceSettingGrammarSlotTypeSetting `field:"optional" json:"grammarSlotTypeSetting" yaml:"grammarSlotTypeSetting"`
}

