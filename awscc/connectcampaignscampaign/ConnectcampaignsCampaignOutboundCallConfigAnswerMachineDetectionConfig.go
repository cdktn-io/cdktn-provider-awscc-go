// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignscampaign


type ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfig struct {
	// Enables detection of prompts (e.g., beep after after a voicemail greeting).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connectcampaigns_campaign#await_answer_machine_prompt ConnectcampaignsCampaign#await_answer_machine_prompt}
	AwaitAnswerMachinePrompt interface{} `field:"optional" json:"awaitAnswerMachinePrompt" yaml:"awaitAnswerMachinePrompt"`
	// Flag to decided whether outbound calls should have answering machine detection enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connectcampaigns_campaign#enable_answer_machine_detection ConnectcampaignsCampaign#enable_answer_machine_detection}
	EnableAnswerMachineDetection interface{} `field:"optional" json:"enableAnswerMachineDetection" yaml:"enableAnswerMachineDetection"`
}

