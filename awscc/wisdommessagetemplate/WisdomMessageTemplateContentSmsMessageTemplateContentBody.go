// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplate


type WisdomMessageTemplateContentSmsMessageTemplateContentBody struct {
	// The container of message template body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_message_template#plain_text WisdomMessageTemplate#plain_text}
	PlainText *WisdomMessageTemplateContentSmsMessageTemplateContentBodyPlainText `field:"optional" json:"plainText" yaml:"plainText"`
}

