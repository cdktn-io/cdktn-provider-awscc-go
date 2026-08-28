// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoiceregistration


type SmsvoiceRegistrationTags struct {
	// The key identifier, or name, of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_registration#key SmsvoiceRegistration#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The string value associated with the key of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_registration#value SmsvoiceRegistration#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

