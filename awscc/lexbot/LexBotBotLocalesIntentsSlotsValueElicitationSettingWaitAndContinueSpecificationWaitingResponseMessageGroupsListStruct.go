// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationWaitingResponseMessageGroupsListStruct struct {
	// The primary message that Amazon Lex should send to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#message LexBot#message}
	Message *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationWaitingResponseMessageGroupsListMessage `field:"optional" json:"message" yaml:"message"`
	// Message variations to send to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#variations LexBot#variations}
	Variations interface{} `field:"optional" json:"variations" yaml:"variations"`
}

