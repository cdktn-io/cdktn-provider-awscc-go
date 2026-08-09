// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightagent


type QuicksightAgentCustomPromptInput struct {
	// Reference to an existing custom prompt profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_agent#existing_prompt QuicksightAgent#existing_prompt}
	ExistingPrompt *QuicksightAgentCustomPromptInputExistingPrompt `field:"optional" json:"existingPrompt" yaml:"existingPrompt"`
	// Parameters for creating a new custom prompt configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_agent#new_prompt QuicksightAgent#new_prompt}
	NewPrompt *QuicksightAgentCustomPromptInputNewPrompt `field:"optional" json:"newPrompt" yaml:"newPrompt"`
}

