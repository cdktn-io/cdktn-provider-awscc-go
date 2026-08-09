// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoicepool


type SmsvoicePoolOptionalKeywords struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/smsvoice_pool#action SmsvoicePool#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/smsvoice_pool#keyword SmsvoicePool#keyword}.
	Keyword *string `field:"optional" json:"keyword" yaml:"keyword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/smsvoice_pool#message SmsvoicePool#message}.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

