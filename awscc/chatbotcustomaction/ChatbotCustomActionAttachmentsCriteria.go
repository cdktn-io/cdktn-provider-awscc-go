// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chatbotcustomaction


type ChatbotCustomActionAttachmentsCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chatbot_custom_action#operator ChatbotCustomAction#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chatbot_custom_action#value ChatbotCustomAction#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chatbot_custom_action#variable_name ChatbotCustomAction#variable_name}.
	VariableName *string `field:"optional" json:"variableName" yaml:"variableName"`
}

