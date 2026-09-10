// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoiceprotectconfiguration


type SmsvoiceProtectConfigurationCountryRuleSetSms struct {
	// The two-letter ISO country code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/smsvoice_protect_configuration#country_code SmsvoiceProtectConfiguration#country_code}
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// The types of protection that can be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/smsvoice_protect_configuration#protect_status SmsvoiceProtectConfiguration#protect_status}
	ProtectStatus *string `field:"optional" json:"protectStatus" yaml:"protectStatus"`
}

