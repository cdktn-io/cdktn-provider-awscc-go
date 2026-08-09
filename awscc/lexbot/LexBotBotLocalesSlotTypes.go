// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesSlotTypes struct {
	// A description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#description LexBot#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Provides information about the external source of the slot type's definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#external_source_setting LexBot#external_source_setting}
	ExternalSourceSetting *LexBotBotLocalesSlotTypesExternalSourceSetting `field:"optional" json:"externalSourceSetting" yaml:"externalSourceSetting"`
	// Unique name for a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#name LexBot#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The built-in slot type used as a parent of this slot type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#parent_slot_type_signature LexBot#parent_slot_type_signature}
	ParentSlotTypeSignature *string `field:"optional" json:"parentSlotTypeSignature" yaml:"parentSlotTypeSignature"`
	// A List of slot type values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#slot_type_values LexBot#slot_type_values}
	SlotTypeValues interface{} `field:"optional" json:"slotTypeValues" yaml:"slotTypeValues"`
	// Contains settings used by Amazon Lex to select a slot value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#value_selection_setting LexBot#value_selection_setting}
	ValueSelectionSetting *LexBotBotLocalesSlotTypesValueSelectionSetting `field:"optional" json:"valueSelectionSetting" yaml:"valueSelectionSetting"`
}

