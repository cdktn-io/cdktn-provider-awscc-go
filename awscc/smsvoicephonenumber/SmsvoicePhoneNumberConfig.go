// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoicephonenumber

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SmsvoicePhoneNumberConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/smsvoice_phone_number#iso_country_code SmsvoicePhoneNumber#iso_country_code}
	IsoCountryCode *string `field:"required" json:"isoCountryCode" yaml:"isoCountryCode"`
	// A keyword is a word that you can search for on a particular phone number or pool.
	//
	// It is also a specific word or phrase that an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number receives a message that begins with a keyword, AWS End User Messaging SMS and Voice responds with a customizable message. Keywords "HELP" and "STOP" are mandatory keywords
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/smsvoice_phone_number#mandatory_keywords SmsvoicePhoneNumber#mandatory_keywords}
	MandatoryKeywords *SmsvoicePhoneNumberMandatoryKeywords `field:"required" json:"mandatoryKeywords" yaml:"mandatoryKeywords"`
	// Indicates if the phone number will be used for text messages, voice messages, or both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/smsvoice_phone_number#number_capabilities SmsvoicePhoneNumber#number_capabilities}
	NumberCapabilities *[]*string `field:"required" json:"numberCapabilities" yaml:"numberCapabilities"`
	// The type of phone number to request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/smsvoice_phone_number#number_type SmsvoicePhoneNumber#number_type}
	NumberType *string `field:"required" json:"numberType" yaml:"numberType"`
	// When set to true the sender ID can't be deleted. By default this is set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/smsvoice_phone_number#deletion_protection_enabled SmsvoicePhoneNumber#deletion_protection_enabled}
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// A keyword is a word that you can search for on a particular phone number or pool.
	//
	// It is also a specific word or phrase that an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number receives a message that begins with a keyword, AWS End User Messaging SMS and Voice responds with a customizable message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/smsvoice_phone_number#optional_keywords SmsvoicePhoneNumber#optional_keywords}
	OptionalKeywords interface{} `field:"optional" json:"optionalKeywords" yaml:"optionalKeywords"`
	// The name of the OptOutList to associate with the phone number. You can use the OptOutListName or OptOutListArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/smsvoice_phone_number#opt_out_list_name SmsvoicePhoneNumber#opt_out_list_name}
	OptOutListName *string `field:"optional" json:"optOutListName" yaml:"optOutListName"`
	// By default this is set to false.
	//
	// When an end recipient sends a message that begins with HELP or STOP to one of your dedicated numbers, AWS End User Messaging SMS and Voice automatically replies with a customizable message and adds the end recipient to the OptOutList. When set to true you're responsible for responding to HELP and STOP requests. You're also responsible for tracking and honoring opt-out requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/smsvoice_phone_number#self_managed_opt_outs_enabled SmsvoicePhoneNumber#self_managed_opt_outs_enabled}
	SelfManagedOptOutsEnabled interface{} `field:"optional" json:"selfManagedOptOutsEnabled" yaml:"selfManagedOptOutsEnabled"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/smsvoice_phone_number#tags SmsvoicePhoneNumber#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// When you set up two-way SMS, you can receive incoming messages from your customers.
	//
	// When one of your customers sends a message to your phone number, the message body is sent to an Amazon SNS topic or Amazon Connect for processing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/smsvoice_phone_number#two_way SmsvoicePhoneNumber#two_way}
	TwoWay *SmsvoicePhoneNumberTwoWay `field:"optional" json:"twoWay" yaml:"twoWay"`
}

