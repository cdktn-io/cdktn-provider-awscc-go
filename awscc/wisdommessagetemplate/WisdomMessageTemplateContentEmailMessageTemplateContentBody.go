// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplate


type WisdomMessageTemplateContentEmailMessageTemplateContentBody struct {
	// The message body, in HTML format, to use in email messages that are based on the message template.
	//
	// We recommend using HTML format for email clients that render HTML content. You can include links, formatted text, and more in an HTML message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wisdom_message_template#html WisdomMessageTemplate#html}
	Html *WisdomMessageTemplateContentEmailMessageTemplateContentBodyHtml `field:"optional" json:"html" yaml:"html"`
	// The message body, in plain text format, to use in email messages that are based on the message template.
	//
	// We recommend using plain text format for email clients that don't render HTML content and clients that are connected to high-latency networks, such as mobile devices.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wisdom_message_template#plain_text WisdomMessageTemplate#plain_text}
	PlainText *WisdomMessageTemplateContentEmailMessageTemplateContentBodyPlainText `field:"optional" json:"plainText" yaml:"plainText"`
}

