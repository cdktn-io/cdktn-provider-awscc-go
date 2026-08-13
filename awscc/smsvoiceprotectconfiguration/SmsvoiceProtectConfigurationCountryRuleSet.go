// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoiceprotectconfiguration


type SmsvoiceProtectConfigurationCountryRuleSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/smsvoice_protect_configuration#mms SmsvoiceProtectConfiguration#mms}.
	Mms interface{} `field:"optional" json:"mms" yaml:"mms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/smsvoice_protect_configuration#sms SmsvoiceProtectConfiguration#sms}.
	Sms interface{} `field:"optional" json:"sms" yaml:"sms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/smsvoice_protect_configuration#voice SmsvoiceProtectConfiguration#voice}.
	Voice interface{} `field:"optional" json:"voice" yaml:"voice"`
}

