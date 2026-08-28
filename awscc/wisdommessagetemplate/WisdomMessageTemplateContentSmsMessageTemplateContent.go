// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplate


type WisdomMessageTemplateContentSmsMessageTemplateContent struct {
	// The body to use in SMS messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#body WisdomMessageTemplate#body}
	Body *WisdomMessageTemplateContentSmsMessageTemplateContentBody `field:"optional" json:"body" yaml:"body"`
}

