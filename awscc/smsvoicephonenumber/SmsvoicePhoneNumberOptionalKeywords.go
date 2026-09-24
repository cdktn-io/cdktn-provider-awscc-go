// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoicephonenumber


type SmsvoicePhoneNumberOptionalKeywords struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/smsvoice_phone_number#action SmsvoicePhoneNumber#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/smsvoice_phone_number#keyword SmsvoicePhoneNumber#keyword}.
	Keyword *string `field:"optional" json:"keyword" yaml:"keyword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/smsvoice_phone_number#message SmsvoicePhoneNumber#message}.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

