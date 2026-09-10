// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chatbotcustomaction


type ChatbotCustomActionAttachments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/chatbot_custom_action#button_text ChatbotCustomAction#button_text}.
	ButtonText *string `field:"optional" json:"buttonText" yaml:"buttonText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/chatbot_custom_action#criteria ChatbotCustomAction#criteria}.
	Criteria interface{} `field:"optional" json:"criteria" yaml:"criteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/chatbot_custom_action#notification_type ChatbotCustomAction#notification_type}.
	NotificationType *string `field:"optional" json:"notificationType" yaml:"notificationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/chatbot_custom_action#variables ChatbotCustomAction#variables}.
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

