// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesVoiceSettings struct {
	// Indicates the type of Amazon Polly voice that Amazon Lex should use for voice interaction with the user.
	//
	// For more information, see the engine parameter of the SynthesizeSpeech operation in the Amazon Polly developer guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lex_bot#engine LexBot#engine}
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The Amazon Polly voice ID that Amazon Lex uses for voice interaction with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lex_bot#voice_id LexBot#voice_id}
	VoiceId *string `field:"optional" json:"voiceId" yaml:"voiceId"`
}

