// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsSlotPriorities struct {
	// The priority that a slot should be elicited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot#priority LexBot#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The name of the slot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot#slot_name LexBot#slot_name}
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
}

