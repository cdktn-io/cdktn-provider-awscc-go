// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplate


type WisdomMessageTemplateContent struct {
	// The content of message template that applies to email channel subtype.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_message_template#email_message_template_content WisdomMessageTemplate#email_message_template_content}
	EmailMessageTemplateContent *WisdomMessageTemplateContentEmailMessageTemplateContent `field:"optional" json:"emailMessageTemplateContent" yaml:"emailMessageTemplateContent"`
	// The content of message template that applies to SMS channel subtype.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_message_template#sms_message_template_content WisdomMessageTemplate#sms_message_template_content}
	SmsMessageTemplateContent *WisdomMessageTemplateContentSmsMessageTemplateContent `field:"optional" json:"smsMessageTemplateContent" yaml:"smsMessageTemplateContent"`
}

