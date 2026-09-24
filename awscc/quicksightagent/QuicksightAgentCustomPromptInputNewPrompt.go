// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightagent


type QuicksightAgentCustomPromptInputNewPrompt struct {
	// Custom instructions for the agent behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_agent#custom_instructions QuicksightAgent#custom_instructions}
	CustomInstructions *string `field:"optional" json:"customInstructions" yaml:"customInstructions"`
	// The identity or persona of the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_agent#identity QuicksightAgent#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// The output style for the agent responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_agent#output_style QuicksightAgent#output_style}
	OutputStyle *string `field:"optional" json:"outputStyle" yaml:"outputStyle"`
	// The desired response length for the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_agent#response_length QuicksightAgent#response_length}
	ResponseLength *string `field:"optional" json:"responseLength" yaml:"responseLength"`
	// The tone used in agent responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_agent#tone QuicksightAgent#tone}
	Tone *string `field:"optional" json:"tone" yaml:"tone"`
}

