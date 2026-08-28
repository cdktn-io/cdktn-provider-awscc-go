// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoicepool

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SmsvoicePoolConfig struct {
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
	// A keyword is a word that you can search for on a particular phone number or pool.
	//
	// It is also a specific word or phrase that an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number receives a message that begins with a keyword, AWS End User Messaging SMS and Voice responds with a customizable message. Keywords "HELP" and "STOP" are mandatory keywords
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_pool#mandatory_keywords SmsvoicePool#mandatory_keywords}
	MandatoryKeywords *SmsvoicePoolMandatoryKeywords `field:"required" json:"mandatoryKeywords" yaml:"mandatoryKeywords"`
	// The origination identity to use such as a PhoneNumberId, PhoneNumberArn, SenderId or SenderIdArn and it's IsoCountryCode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_pool#origination_identities SmsvoicePool#origination_identities}
	OriginationIdentities *[]*string `field:"required" json:"originationIdentities" yaml:"originationIdentities"`
	// When set to true the pool can't be deleted. By default this is set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_pool#deletion_protection_enabled SmsvoicePool#deletion_protection_enabled}
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// A keyword is a word that you can search for on a particular phone number or pool.
	//
	// It is also a specific word or phrase that an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number receives a message that begins with a keyword, AWS End User Messaging SMS and Voice responds with a customizable message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_pool#optional_keywords SmsvoicePool#optional_keywords}
	OptionalKeywords interface{} `field:"optional" json:"optionalKeywords" yaml:"optionalKeywords"`
	// The name of the OptOutList to associate with the pool. You can use the OptOutListName or OptOutListArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_pool#opt_out_list_name SmsvoicePool#opt_out_list_name}
	OptOutListName *string `field:"optional" json:"optOutListName" yaml:"optOutListName"`
	// By default this is set to false.
	//
	// When an end recipient sends a message that begins with HELP or STOP to one of your dedicated numbers, AWS End User Messaging SMS and Voice automatically replies with a customizable message and adds the end recipient to the OptOutList. When set to true you're responsible for responding to HELP and STOP requests. You're also responsible for tracking and honoring opt-out requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_pool#self_managed_opt_outs_enabled SmsvoicePool#self_managed_opt_outs_enabled}
	SelfManagedOptOutsEnabled interface{} `field:"optional" json:"selfManagedOptOutsEnabled" yaml:"selfManagedOptOutsEnabled"`
	// Indicates whether shared routes are enabled for the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_pool#shared_routes_enabled SmsvoicePool#shared_routes_enabled}
	SharedRoutesEnabled interface{} `field:"optional" json:"sharedRoutesEnabled" yaml:"sharedRoutesEnabled"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_pool#tags SmsvoicePool#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// When you set up two-way SMS, you can receive incoming messages from your customers.
	//
	// When one of your customers sends a message to your phone number, the message body is sent to an Amazon SNS topic or Amazon Connect for processing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_pool#two_way SmsvoicePool#two_way}
	TwoWay *SmsvoicePoolTwoWay `field:"optional" json:"twoWay" yaml:"twoWay"`
}

