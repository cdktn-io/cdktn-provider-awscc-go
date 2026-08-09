// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplate


type WisdomMessageTemplateContentEmailMessageTemplateContent struct {
	// The body to use in email messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_message_template#body WisdomMessageTemplate#body}
	Body *WisdomMessageTemplateContentEmailMessageTemplateContentBody `field:"optional" json:"body" yaml:"body"`
	// The email headers to include in email messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_message_template#headers WisdomMessageTemplate#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// The subject line, or title, to use in email messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_message_template#subject WisdomMessageTemplate#subject}
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

