// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicCustomInstructions struct {
	// <p>A text field for providing additional guidance or context for response generation.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_topic#custom_instructions_string QuicksightTopic#custom_instructions_string}
	CustomInstructionsString *string `field:"optional" json:"customInstructionsString" yaml:"customInstructionsString"`
}

