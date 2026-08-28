// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoiceconfigurationset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SmsvoiceConfigurationSetConfig struct {
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
	// The name to use for the configuration set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_configuration_set#configuration_set_name SmsvoiceConfigurationSet#configuration_set_name}
	ConfigurationSetName *string `field:"optional" json:"configurationSetName" yaml:"configurationSetName"`
	// The default sender ID to set for the ConfigurationSet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_configuration_set#default_sender_id SmsvoiceConfigurationSet#default_sender_id}
	DefaultSenderId *string `field:"optional" json:"defaultSenderId" yaml:"defaultSenderId"`
	// An event destination is a location where you send message events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_configuration_set#event_destinations SmsvoiceConfigurationSet#event_destinations}
	EventDestinations interface{} `field:"optional" json:"eventDestinations" yaml:"eventDestinations"`
	// Set to true to enable message feedback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_configuration_set#message_feedback_enabled SmsvoiceConfigurationSet#message_feedback_enabled}
	MessageFeedbackEnabled interface{} `field:"optional" json:"messageFeedbackEnabled" yaml:"messageFeedbackEnabled"`
	// The unique identifier for the protect configuration to be associated to the configuration set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_configuration_set#protect_configuration_id SmsvoiceConfigurationSet#protect_configuration_id}
	ProtectConfigurationId *string `field:"optional" json:"protectConfigurationId" yaml:"protectConfigurationId"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/smsvoice_configuration_set#tags SmsvoiceConfigurationSet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

