// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoicephonenumber


type SmsvoicePhoneNumberMandatoryKeywords struct {
	// A keyword is a word that you can search for on a particular phone number or pool.
	//
	// It is also a specific word or phrase that an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number receives a message that begins with a keyword, AWS End User Messaging SMS and Voice responds with a customizable message. Keywords "HELP" and "STOP" are mandatory keywords
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_phone_number#help SmsvoicePhoneNumber#help}
	Help *SmsvoicePhoneNumberMandatoryKeywordsHelp `field:"required" json:"help" yaml:"help"`
	// A keyword is a word that you can search for on a particular phone number or pool.
	//
	// It is also a specific word or phrase that an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number receives a message that begins with a keyword, AWS End User Messaging SMS and Voice responds with a customizable message. Keywords "HELP" and "STOP" are mandatory keywords
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_phone_number#stop SmsvoicePhoneNumber#stop}
	Stop *SmsvoicePhoneNumberMandatoryKeywordsStop `field:"required" json:"stop" yaml:"stop"`
}

