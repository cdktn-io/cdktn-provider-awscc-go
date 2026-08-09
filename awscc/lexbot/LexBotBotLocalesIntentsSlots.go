// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsSlots struct {
	// A description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#description LexBot#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether a slot can return multiple values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#multiple_values_setting LexBot#multiple_values_setting}
	MultipleValuesSetting *LexBotBotLocalesIntentsSlotsMultipleValuesSetting `field:"optional" json:"multipleValuesSetting" yaml:"multipleValuesSetting"`
	// Unique name for a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#name LexBot#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Determines whether Amazon Lex obscures slot values in conversation logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#obfuscation_setting LexBot#obfuscation_setting}
	ObfuscationSetting *LexBotBotLocalesIntentsSlotsObfuscationSetting `field:"optional" json:"obfuscationSetting" yaml:"obfuscationSetting"`
	// The slot type name that is used in the slot. Allows for custom and built-in slot type names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#slot_type_name LexBot#slot_type_name}
	SlotTypeName *string `field:"optional" json:"slotTypeName" yaml:"slotTypeName"`
	// Settings that you can use for eliciting a slot value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#value_elicitation_setting LexBot#value_elicitation_setting}
	ValueElicitationSetting *LexBotBotLocalesIntentsSlotsValueElicitationSetting `field:"optional" json:"valueElicitationSetting" yaml:"valueElicitationSetting"`
}

