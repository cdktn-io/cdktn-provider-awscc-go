// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chatbotcustomaction


type ChatbotCustomActionDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chatbot_custom_action#command_text ChatbotCustomAction#command_text}.
	CommandText *string `field:"required" json:"commandText" yaml:"commandText"`
}

