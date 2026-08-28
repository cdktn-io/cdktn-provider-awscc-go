// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesSlotTypesValueSelectionSettingAdvancedRecognitionSetting struct {
	// Enables using slot values as a custom vocabulary when recognizing user utterances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#audio_recognition_strategy LexBot#audio_recognition_strategy}
	AudioRecognitionStrategy *string `field:"optional" json:"audioRecognitionStrategy" yaml:"audioRecognitionStrategy"`
}

